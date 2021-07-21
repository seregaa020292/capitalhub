//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package session

import (
	"context"
	"github.com/google/uuid"
	"github.com/seregaa020292/capitalhub/internal/auth/model"
)

// Session repository
type SessRepository interface {
	CreateSession(ctx context.Context, session *model.Session, expire int) (string, error)
	CleanMaxSession(ctx context.Context, userID uuid.UUID) int64
	GetSessionByID(ctx context.Context, userID uuid.UUID, sessionID string) (*model.Session, error)
	RefreshByID(ctx context.Context, sess *model.Session, newSessionID string, expire int) (string, error)
	DeleteByID(ctx context.Context, userID uuid.UUID, sessionID string) error
}
