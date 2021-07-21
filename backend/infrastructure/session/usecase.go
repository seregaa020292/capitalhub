//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package session

import (
	"context"
	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/models"
)

// Session use case
type UseCase interface {
	CreateSession(ctx context.Context, session *models.Session, expire int) (string, error)
	CleanMaxSession(ctx context.Context, userID uuid.UUID) int64
	GetSessionByID(ctx context.Context, userID uuid.UUID, sessionID string) (*models.Session, error)
	RefreshByID(ctx context.Context, session *models.Session, newSessionID string, expire int) (string, error)
	DeleteByID(ctx context.Context, userID uuid.UUID, sessionID string) error
}
