//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package auth

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/auth/model"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Auth repository interface
type Repository interface {
	Register(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Confirmed(ctx context.Context, code uuid.UUID) error
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*model.User, error)
	FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*model.UsersList, error)
	FindByEmail(ctx context.Context, user *model.User) (*model.User, error)
	GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*model.UsersList, error)
}
