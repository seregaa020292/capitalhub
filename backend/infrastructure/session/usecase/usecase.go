package usecase

import (
	"context"
	"github.com/google/uuid"

	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/session"
	"github.com/seregaa020292/capitalhub/internal/models"
)

// Session use case
type sessionUC struct {
	sessionRepo session.SessRepository
	cfg         *config.Config
}

// New session use case constructor
func NewSessionUseCase(sessionRepo session.SessRepository, cfg *config.Config) session.UseCase {
	return &sessionUC{sessionRepo: sessionRepo, cfg: cfg}
}

// Создаем новую сессию
func (u *sessionUC) CreateSession(ctx context.Context, session *models.Session, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.CreateSession")
	defer span.Finish()

	return u.sessionRepo.CreateSession(ctx, session, expire)
}

// Очищаем старые сессии если превышен лимит
func (u *sessionUC) CleanMaxSession(ctx context.Context, userID uuid.UUID) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.CleanMaxSession")
	defer span.Finish()

	return u.sessionRepo.CleanMaxSession(ctx, userID)
}

// Обновляем сессию по id
func (u *sessionUC) RefreshByID(ctx context.Context, sess *models.Session, newSessionID string, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.RefreshByID")
	defer span.Finish()

	return u.sessionRepo.RefreshByID(ctx, sess, newSessionID, expire)
}

// Удаляем сессию по id
func (u *sessionUC) DeleteByID(ctx context.Context, userID uuid.UUID, sessionID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.DeleteByID")
	defer span.Finish()

	return u.sessionRepo.DeleteByID(ctx, userID, sessionID)
}

// Получаем сессию по id
func (u *sessionUC) GetSessionByID(ctx context.Context, userID uuid.UUID, sessionID string) (*models.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionUC.GetSessionByID")
	defer span.Finish()

	return u.sessionRepo.GetSessionByID(ctx, userID, sessionID)
}
