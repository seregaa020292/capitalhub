package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/auth"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Auth Repository
type authRepo struct {
	db *sqlx.DB
}

// Auth Repository constructor
func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepo{db: db}
}

// Создание пользователя
func (repo *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	u := &models.User{}
	if err := repo.db.QueryRowxContext(ctx, createUserQuery,
		&user.Name, &user.Email, &user.Password, &user.Role, &user.Avatar,
	).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.StructScan")
	}

	return u, nil
}

// Подтверждение почты пользователя, установив поле confirmed = NULL
func (repo *authRepo) Confirmed(ctx context.Context, code uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Confirmed")
	defer span.Finish()

	result, err := repo.db.ExecContext(ctx, confirmedUserQuery, code)
	if err != nil {
		return errors.Wrap(err, "authRepo.Confirmed.GetContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "authRepo.Confirmed.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "authRepo.Confirmed.rowsAffected")
	}

	return nil
}

// Изменение данных пользователя
func (repo *authRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Update")
	defer span.Finish()

	u := &models.User{}
	if err := repo.db.GetContext(ctx, u, updateUserQuery, &user.Name, &user.Email,
		&user.Role, &user.Avatar, &user.UserID,
	); err != nil {
		return nil, errors.Wrap(err, "authRepo.Update.GetContext")
	}

	return u, nil
}

// Удаление пользователя
func (repo *authRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Delete")
	defer span.Finish()

	result, err := repo.db.ExecContext(ctx, deleteUserQuery, userID)
	if err != nil {
		return errors.WithMessage(err, "authRepo Delete ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "authRepo.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "authRepo.Delete.rowsAffected")
	}

	return nil
}

// Вернуть пользователя по ид
func (repo *authRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.GetByID")
	defer span.Finish()

	user := &models.User{}
	if err := repo.db.QueryRowxContext(ctx, getUserQuery, userID).StructScan(user); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetByID.QueryRowxContext")
	}
	return user, nil
}

// Найти пользователей по имени
func (repo *authRepo) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.FindByName")
	defer span.Finish()

	var totalCount int
	if err := repo.db.GetContext(ctx, &totalCount, getTotalCount, name); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByName.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.UsersList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Users:      make([]*models.User, 0),
		}, nil
	}

	rows, err := repo.db.QueryxContext(ctx, findUsers, name, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByName.QueryxContext")
	}
	defer rows.Close()

	var users = make([]*models.User, 0, query.GetSize())
	for rows.Next() {
		var user models.User
		if err = rows.StructScan(&user); err != nil {
			return nil, errors.Wrap(err, "authRepo.FindByName.StructScan")
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByName.rows.Err")
	}

	return &models.UsersList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Users:      users,
	}, nil
}

// Вернуть всех пользователей
func (repo *authRepo) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.GetUsers")
	defer span.Finish()

	var totalCount int
	if err := repo.db.GetContext(ctx, &totalCount, getTotal); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetUsers.GetContext.totalCount")
	}

	if totalCount == 0 {
		return &models.UsersList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			Users:      make([]*models.User, 0),
		}, nil
	}

	var users = make([]*models.User, 0, pq.GetSize())
	if err := repo.db.SelectContext(
		ctx,
		&users,
		getUsers,
		pq.GetOrderBy(),
		pq.GetOffset(),
		pq.GetLimit(),
	); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetUsers.SelectContext")
	}

	return &models.UsersList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
		Page:       pq.GetPage(),
		Size:       pq.GetSize(),
		HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		Users:      users,
	}, nil
}

// Найти пользователя по email
func (repo *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.FindByEmail")
	defer span.Finish()

	foundUser := &models.User{}
	if err := repo.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundUser); err != nil {
		return nil, errors.Wrap(err, "authRepo.FindByEmail.QueryRowxContext")
	}
	return foundUser, nil
}
