//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package auth

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/auth/model"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Auth useCase interface
type UseCase interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	Login(ctx context.Context, user *model.User) (*model.UserWithToken, error)
	Confirmed(ctx context.Context, code uuid.UUID) error
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	GetRefreshByID(ctx context.Context, userID uuid.UUID) (*model.UserWithToken, error)
	FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*model.UsersList, error)
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*model.UsersList, error)
	UploadAvatar(ctx context.Context, userID uuid.UUID, file model.UploadInput) (*model.User, error)
}
