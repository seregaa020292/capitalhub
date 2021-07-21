package http

import (
	"bytes"
	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/service"
	"github.com/seregaa020292/capitalhub/infrastructure/session"
	"github.com/seregaa020292/capitalhub/internal/auth"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
	"github.com/seregaa020292/capitalhub/pkg/csrf"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Auth handlers
type authHandlers struct {
	cfg              *config.Config
	authUseCase      auth.UseCase
	portfolioUseCase portfolio.UseCase
	sessUseCase      session.UseCase
	emailService     service.Email
	logger           logger.Logger
}

// NewAuthHandlers Auth handlers constructor
func NewAuthHandlers(
	cfg *config.Config,
	authUseCase auth.UseCase,
	portfolioUseCase portfolio.UseCase,
	sessUseCase session.UseCase,
	emailService service.Email,
	log logger.Logger,
) auth.Handlers {
	return &authHandlers{
		cfg:              cfg,
		authUseCase:      authUseCase,
		portfolioUseCase: portfolioUseCase,
		sessUseCase:      sessUseCase,
		emailService:     emailService,
		logger:           log,
	}
}

// Register godoc
// @Summary Зарегистрировать нового пользователя
// @Description Зарегистрировать и вернуть пользователя и токен
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body models.User true "register user"
// @Success 201
// @Router /auth/register [post]
func (handler *authHandlers) Register() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.Register")
		defer span.Finish()

		user := &models.User{}
		if err := utils.ReadRequest(echoCtx, user); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		createdUser, err := handler.authUseCase.Register(ctx, user)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		if _, err := handler.portfolioUseCase.CreateFirst(ctx, createdUser.UserID); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		handler.emailService.SendConfirmedMail(createdUser.Email, *createdUser.Confirmed)

		return echoCtx.NoContent(http.StatusCreated)
	}
}

// Confirmed godoc
// @Summary Подтверждение почты
// @Description Подтверждается почта
// @Tags Auth
// @Accept json
// @Produce json
// @Param code query string true "code" Format(code)
// @Success 200
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/confirmed [get]
func (handler *authHandlers) Confirmed() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.Confirmed")
		defer span.Finish()

		codeID, err := uuid.Parse(echoCtx.QueryParam("code"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		if err := handler.authUseCase.Confirmed(ctx, codeID); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.NoContent(http.StatusOK)
	}
}

// Login godoc
// @Summary Вход пользователя
// @Description Войти в систему, вернуть пользователя и установить сеанс
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.UserWithToken
// @Router /auth/login [post]
func (handler *authHandlers) Login() echo.HandlerFunc {
	type Login struct {
		Fingerprint string `json:"fingerprint"`
		Email       string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password    string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.Login")
		defer span.Finish()

		login := &Login{}
		if err := utils.ReadRequest(echoCtx, login); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		userWithToken, err := handler.authUseCase.Login(ctx, &models.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		// Записываем refresh токен в хранилище
		if _, err := handler.sessUseCase.CreateSession(ctx, &models.Session{
			Fingerprint: login.Fingerprint,
			SessionID:   userWithToken.RefreshToken.Token,
			UserID:      userWithToken.User.UserID,
		}, handler.cfg.Cookie.MaxAge); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		// Очищаем старые сессии если превышен лимит
		handler.sessUseCase.CleanMaxSession(ctx, userWithToken.User.UserID)

		// Записываем session refresh token в cookie
		echoCtx.SetCookie(utils.SetRefreshTokenCookie(handler.cfg, userWithToken.RefreshToken.Token))

		// Создаем и отдаем в заголовке CSRF токен
		utils.SetCSRFHeader(echoCtx, userWithToken.AccessToken.Token, handler.cfg.Server.CsrfSalt, handler.logger)

		return echoCtx.JSON(http.StatusOK, userWithToken)
	}
}

// CheckLogged godoc
// @Summary Проверить зарегистрированного пользователя
// @Description Проверить токен на действительность
// @Security Auth
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.UserBase
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/check [get]
func (handler *authHandlers) CheckLogged() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, _ := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.checkLogged")
		defer span.Finish()

		user, ok := echoCtx.Get("user").(*models.User)
		if !ok {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, httpErrors.NewUnauthorizedError(httpErrors.Unauthorized))
		}

		// Создаем и отдаем в заголовке CSRF токен
		sid, _ := echoCtx.Get("sid").(string)
		utils.SetCSRFHeader(echoCtx, sid, handler.cfg.Server.CsrfSalt, handler.logger)

		return echoCtx.JSON(http.StatusOK, &models.UserBase{User: user})
	}
}

// RefreshToken godoc
// @Summary Обновить токен
// @Description Обновить и вернуть новые токены
// @Security Auth
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.Tokens
// @Router /auth/refresh [post]
func (handler *authHandlers) RefreshToken() echo.HandlerFunc {
	type Refresh struct {
		Fingerprint string `json:"fingerprint"`
	}
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.refreshToken")
		defer span.Finish()

		refresh := &Refresh{}
		if err := utils.ReadRequest(echoCtx, refresh); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		// Парсим токен из заголовка Authorization
		claims := jwt.MapClaims{}
		if _, err := utils.ParseJWT(echoCtx, handler.cfg, claims); err != nil {
			if !errors.Is(err, httpErrors.ExpiredJWTToken) {
				utils.LogResponseError(echoCtx, handler.logger, err)
				return echoCtx.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(err))
			}
		}
		// Получаем ид пользователя
		userUUID, err := utils.ParseUserIDFromJWT(claims["id"])
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		// Получаем токен из Cookie
		cookie, err := echoCtx.Cookie(handler.cfg.Auth.NameRefreshToken)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				utils.LogResponseError(echoCtx, handler.logger, err)
				return echoCtx.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(err))
			}
			utils.LogResponseError(echoCtx, handler.logger, err)
			return echoCtx.JSON(http.StatusInternalServerError, httpErrors.NewInternalServerError(err))
		}

		// Получаем и проверяем токен из хранилища
		sessCurrent, err := handler.sessUseCase.GetSessionByID(ctx, userUUID, cookie.Value)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, httpErrors.NewUnauthorizedError(err))
		}
		if sessCurrent.Fingerprint != refresh.Fingerprint {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, httpErrors.NewUnauthorizedError(httpErrors.FingerprintNotMatch))
		}

		// Получаем пользователя из БД и генерируем токены
		userWithToken, err := handler.authUseCase.GetRefreshByID(echoCtx.Request().Context(), userUUID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		// Обновляем токен в хранилище
		if _, err := handler.sessUseCase.RefreshByID(ctx, &models.Session{
			SessionID:   cookie.Value,
			UserID:      userUUID,
			Fingerprint: refresh.Fingerprint,
		}, userWithToken.RefreshToken.Token, handler.cfg.Cookie.MaxAge); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		// Записываем session refresh token в cookie
		echoCtx.SetCookie(utils.SetRefreshTokenCookie(handler.cfg, userWithToken.RefreshToken.Token))

		// Создаем и отдаем в заголовке CSRF токен
		utils.SetCSRFHeader(echoCtx, userWithToken.AccessToken.Token, handler.cfg.Server.CsrfSalt, handler.logger)

		return echoCtx.JSON(http.StatusOK, &models.Tokens{
			AccessToken:  userWithToken.AccessToken,
			RefreshToken: userWithToken.RefreshToken,
		})
	}
}

// Logout godoc
// @Summary Выйти из системы
// @Description Удаление сеанса пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string	"ok"
// @Router /auth/logout [post]
func (handler *authHandlers) Logout() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.Logout")
		defer span.Finish()

		cookie, err := echoCtx.Cookie(handler.cfg.Auth.NameRefreshToken)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				utils.LogResponseError(echoCtx, handler.logger, err)
				return echoCtx.JSON(http.StatusUnauthorized, httpErrors.NewUnauthorizedError(err))
			}
			utils.LogResponseError(echoCtx, handler.logger, err)
			return echoCtx.JSON(http.StatusInternalServerError, httpErrors.NewInternalServerError(err))
		}

		claims, err := utils.ExtractJWTFromRequest(echoCtx, handler.cfg)
		if err == nil {
			userUUID, err := utils.ParseUserIDFromJWT(claims["id"])
			if err == nil {
				handler.sessUseCase.DeleteByID(ctx, userUUID, cookie.Value)
			}
		}

		utils.DeleteSessionCookie(echoCtx, handler.cfg.Auth.NameRefreshToken)

		return echoCtx.NoContent(http.StatusOK)
	}
}

// Update godoc
// @Summary Обновить пользователя
// @Description Обновить существующего пользователя
// @Security Auth
// @Tags Auth
// @Accept json
// @Param id path int true "user_id"
// @Produce json
// @Success 200 {object} models.User
// @Router /auth/{id} [put]
func (handler *authHandlers) Update() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.Update")
		defer span.Finish()

		uID, err := uuid.Parse(echoCtx.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		user := &models.User{}
		user.UserID = uID

		if err = utils.ReadRequest(echoCtx, user); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		updatedUser, err := handler.authUseCase.Update(ctx, user)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, updatedUser)
	}
}

// GetUserByID godoc
// @Summary Получить пользователя по идентификатору
// @Tags Auth
// @Accept json
// @Produce json
// @Param id path int true "user_id"
// @Success 200 {object} models.User
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/{id} [get]
func (handler *authHandlers) GetUserByID() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.GetUserByID")
		defer span.Finish()

		uID, err := uuid.Parse(echoCtx.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		user, err := handler.authUseCase.GetByID(ctx, uID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, user)
	}
}

// Delete
// @Summary Удалить учетную запись пользователя
// @Security Auth
// @Tags Auth
// @Accept json
// @Param id path int true "user_id"
// @Produce json
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/{id} [delete]
func (handler *authHandlers) Delete() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.Delete")
		defer span.Finish()

		uID, err := uuid.Parse(echoCtx.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		if err = handler.authUseCase.Delete(ctx, uID); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.NoContent(http.StatusOK)
	}
}

// FindByName godoc
// @Summary Найти пользователя по имени
// @Tags Auth
// @Accept json
// @Param name query string false "username" Format(username)
// @Produce json
// @Success 200 {object} models.UsersList
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/find [get]
func (handler *authHandlers) FindByName() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.FindByName")
		defer span.Finish()

		if echoCtx.QueryParam("name") == "" {
			utils.LogResponseError(echoCtx, handler.logger, httpErrors.NewBadRequestError("name is required"))
			return echoCtx.JSON(http.StatusBadRequest, httpErrors.NewBadRequestError("name is required"))
		}

		paginationQuery, err := utils.GetPaginationFromCtx(echoCtx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		response, err := handler.authUseCase.FindByName(ctx, echoCtx.QueryParam("name"), paginationQuery)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, response)
	}
}

// GetUsers godoc
// @Summary Получить список всех пользователей
// @Tags Auth
// @Accept json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Produce json
// @Success 200 {object} models.UsersList
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/all [get]
func (handler *authHandlers) GetUsers() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.GetUsers")
		defer span.Finish()

		paginationQuery, err := utils.GetPaginationFromCtx(echoCtx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		usersList, err := handler.authUseCase.GetUsers(ctx, paginationQuery)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, usersList)
	}
}

// GetCSRFToken godoc
// @Summary Получить токен CSRF
// @Description Получить токен CSRF, обязательный cookie сеанса авторизации
// @Security Auth
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {string} string "Ok"
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/token [get]
func (handler *authHandlers) GetCSRFToken() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, _ := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.GetCSRFToken")
		defer span.Finish()

		sid, ok := echoCtx.Get("sid").(string)
		if !ok {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, httpErrors.NewUnauthorizedError(httpErrors.Unauthorized))
		}
		token := csrf.MakeToken(sid, handler.cfg.Server.CsrfSalt, handler.logger)
		echoCtx.Response().Header().Set(csrf.CSRFHeader, token)
		echoCtx.Response().Header().Set("Access-Control-Expose-Headers", csrf.CSRFHeader)

		return echoCtx.NoContent(http.StatusOK)
	}
}

// UploadAvatar godoc
// @Summary Опубликовать аватар пользователя
// @Security Auth
// @Tags Auth
// @Accept json
// @Produce json
// @Param file formData file true "Body with image file"
// @Param bucket query string true "aws s3 bucket" Format(bucket)
// @Param id path int true "user_id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestError
// @Router /auth/{id}/avatar [post]
func (handler *authHandlers) UploadAvatar() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "authHandlers.UploadAvatar")
		defer span.Finish()

		bucket := echoCtx.QueryParam("bucket")
		uID, err := uuid.Parse(echoCtx.Param("user_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		image, err := utils.ReadImage(echoCtx, "file")
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		file, err := image.Open()
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}
		defer file.Close()

		binaryImage := bytes.NewBuffer(nil)
		if _, err = io.Copy(binaryImage, file); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		contentType, err := utils.CheckImageFileContentType(binaryImage.Bytes())
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		reader := bytes.NewReader(binaryImage.Bytes())

		updatedUser, err := handler.authUseCase.UploadAvatar(ctx, uID, models.UploadInput{
			File:        reader,
			Name:        image.Filename,
			Size:        image.Size,
			ContentType: contentType,
			BucketName:  bucket,
		})
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, updatedUser)
	}
}
