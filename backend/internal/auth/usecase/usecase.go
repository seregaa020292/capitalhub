package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/auth"
	"github.com/seregaa020292/capitalhub/internal/auth/model"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

const (
	basePrefix    = "api-auth:"
	cacheDuration = 3600
)

// Auth UseCase
type authUC struct {
	cfg       *config.Config
	authRepo  auth.Repository
	redisRepo auth.RedisRepository
	awsRepo   auth.AWSRepository
	log       logger.Logger
}

// Auth UseCase constructor
func NewAuthUseCase(
	cfg *config.Config,
	authRepo auth.Repository,
	redisRepo auth.RedisRepository,
	awsRepo auth.AWSRepository,
	log logger.Logger,
) auth.UseCase {
	return &authUC{cfg: cfg, authRepo: authRepo, redisRepo: redisRepo, awsRepo: awsRepo, log: log}
}

// Создаем нового пользователя
func (useCase *authUC) Register(ctx context.Context, user *model.User) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Register")
	defer span.Finish()

	existsUser, err := useCase.authRepo.FindByEmail(ctx, user)
	if existsUser != nil || err == nil {
		return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailAlreadyExists, nil)
	}

	if err = user.PrepareCreate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "authUC.Register.PrepareCreate"))
	}

	createdUser, err := useCase.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	createdUser.SanitizePassword()

	return createdUser, nil
}

// Подтверждение почты пользователя
func (useCase *authUC) Confirmed(ctx context.Context, code uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Confirmed")
	defer span.Finish()

	return useCase.authRepo.Confirmed(ctx, code)
}

// Авторизация пользователя, возвращает модель пользователя с токеном jwt, refresh
func (useCase *authUC) Login(ctx context.Context, user *model.User) (*model.UserWithToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Login")
	defer span.Finish()

	// Есть ли пользователь с email
	foundUser, err := useCase.authRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, err
	}

	// Подтвержден ли аккаунт
	if !foundUser.IsConfirmed() {
		return nil, httpErrors.NewRestErrorWithMessage(http.StatusBadRequest, httpErrors.ErrEmailNotConfirmed, nil)
	}

	// Валиден пароль
	if err = foundUser.ComparePasswords(user.Password); err != nil {
		return nil, httpErrors.NewUnauthorizedError(errors.Wrap(err, "authUC.GetUsers.ComparePasswords"))
	}
	foundUser.SanitizePassword()

	// Создаем токены access & refresh
	token, err := utils.GenerateTokens(foundUser, useCase.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.GetUsers.GenerateTokens"))
	}

	return &model.UserWithToken{
		User: foundUser,
		AccessToken: &model.AccessToken{
			Token:       token.AccessToken,
			PrefixToken: useCase.cfg.Auth.PrefixAccessToken,
		},
		RefreshToken: &model.RefreshToken{
			Token: token.RefreshToken,
		},
	}, nil
}

// Обновить существующего пользователя
func (useCase *authUC) Update(ctx context.Context, user *model.User) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Update")
	defer span.Finish()

	if err := user.PrepareUpdate(); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.Wrap(err, "authUC.Register.PrepareUpdate"))
	}

	updatedUser, err := useCase.authRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	updatedUser.SanitizePassword()

	if err = useCase.redisRepo.DeleteUserCtx(ctx, useCase.generateUserKey(user.UserID.String())); err != nil {
		useCase.log.Errorf("AuthUC.Update.DeleteUserCtx: %s", err)
	}

	updatedUser.SanitizePassword()

	return updatedUser, nil
}

// Удаляем пользователя
func (useCase *authUC) Delete(ctx context.Context, userID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.Delete")
	defer span.Finish()

	if err := useCase.authRepo.Delete(ctx, userID); err != nil {
		return err
	}

	if err := useCase.redisRepo.DeleteUserCtx(ctx, useCase.generateUserKey(userID.String())); err != nil {
		useCase.log.Errorf("AuthUC.Delete.DeleteUserCtx: %s", err)
	}

	return nil
}

// Возвращаем пользователя с новыми токенами
func (useCase *authUC) GetRefreshByID(ctx context.Context, userID uuid.UUID) (*model.UserWithToken, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.GetRefreshByID")
	defer span.Finish()

	user, err := useCase.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Создаем токены access & refresh
	token, err := utils.GenerateTokens(user, useCase.cfg)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.GetUsers.GenerateTokens"))
	}

	return &model.UserWithToken{
		User: user,
		AccessToken: &model.AccessToken{
			Token:       token.AccessToken,
			PrefixToken: useCase.cfg.Auth.PrefixAccessToken,
		},
		RefreshToken: &model.RefreshToken{
			Token: token.RefreshToken,
		},
	}, nil
}

// Получаем пользователя по id
func (useCase *authUC) GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.GetByID")
	defer span.Finish()

	cachedUser, err := useCase.redisRepo.GetByIDCtx(ctx, useCase.generateUserKey(userID.String()))
	if err != nil {
		useCase.log.Errorf("authUC.GetByID.GetByIDCtx: %v", err)
	}
	if cachedUser != nil {
		return cachedUser, nil
	}

	user, err := useCase.authRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err = useCase.redisRepo.SetUserCtx(ctx, useCase.generateUserKey(userID.String()), cacheDuration, user); err != nil {
		useCase.log.Errorf("authUC.GetByID.SetUserCtx: %v", err)
	}

	user.SanitizePassword()

	return user, nil
}

// Находим пользователей по имени
func (useCase *authUC) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*model.UsersList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.FindByName")
	defer span.Finish()

	return useCase.authRepo.FindByName(ctx, name, query)
}

// Получение пользователей с разбивкой на страницы
func (useCase *authUC) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*model.UsersList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.GetUsers")
	defer span.Finish()

	return useCase.authRepo.GetUsers(ctx, pq)
}

// Загрузить аватар пользователя
func (useCase *authUC) UploadAvatar(ctx context.Context, userID uuid.UUID, file model.UploadInput) (*model.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authUC.UploadAvatar")
	defer span.Finish()

	uploadInfo, err := useCase.awsRepo.PutObject(ctx, file)
	if err != nil {
		return nil, httpErrors.NewInternalServerError(errors.Wrap(err, "authUC.UploadAvatar.PutObject"))
	}

	avatarURL := useCase.generateAWSMinioURL(file.BucketName, uploadInfo.Key)

	updatedUser, err := useCase.authRepo.Update(ctx, &model.User{
		UserID: userID,
		Avatar: &avatarURL,
	})
	if err != nil {
		return nil, err
	}

	updatedUser.SanitizePassword()

	return updatedUser, nil
}

func (useCase *authUC) generateUserKey(userID string) string {
	return fmt.Sprintf("%s:%s", basePrefix, userID)
}

func (useCase *authUC) generateAWSMinioURL(bucket string, key string) string {
	return fmt.Sprintf("%s/minio/%s/%s", useCase.cfg.AWS.MinioEndpoint, bucket, key)
}
