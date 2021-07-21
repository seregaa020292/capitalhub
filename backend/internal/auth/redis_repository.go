//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package auth

import (
	"context"
	"github.com/seregaa020292/capitalhub/internal/auth/model"
)

// Auth Redis repository interface
type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*model.User, error)
	SetUserCtx(ctx context.Context, key string, seconds int, user *model.User) error
	DeleteUserCtx(ctx context.Context, key string) error
}
