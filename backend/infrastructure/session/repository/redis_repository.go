package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/infrastructure/session"
	"github.com/seregaa020292/capitalhub/internal/auth/model"
)

const (
	basePrefix = "api-session:"
)

// Session repository
type sessionRepo struct {
	redisClient *redis.Client
	basePrefix  string
	cfg         *config.Config
}

// Session repository constructor
func NewSessionRepository(redisClient *redis.Client, cfg *config.Config) session.SessRepository {
	return &sessionRepo{redisClient: redisClient, basePrefix: basePrefix, cfg: cfg}
}

// Создаем сессию в Redis
func (repository *sessionRepo) CreateSession(ctx context.Context, sess *model.Session, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.CreateSession")
	defer span.Finish()

	sessionKey := repository.createKey(sess.UserID, sess.SessionID)

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		return "", errors.WithMessage(err, "sessionRepo.CreateSession.json.Marshal")
	}
	expireTime := time.Second * time.Duration(expire)
	if err = repository.redisClient.Set(ctx, sessionKey, sessBytes, expireTime).Err(); err != nil {
		return "", errors.Wrap(err, "sessionRepo.CreateSession.redisClient.Set")
	}
	return sessionKey, nil
}

// Максимум до 5 одновременных рефреш-сессий для пользователя
func (repository *sessionRepo) CleanMaxSession(ctx context.Context, userID uuid.UUID) int64 {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.CleanMaxSession")
	defer span.Finish()

	sessionList := repository.redisClient.Keys(ctx, repository.createKey(userID, "*")).Val()

	if len(sessionList) <= repository.cfg.Auth.MaxRefreshSession {
		return 0
	}

	type sessionTTL struct {
		key string
		exp int64
	}
	sessionTTLList := make([]sessionTTL, 0, len(sessionList))
	for _, sessKey := range sessionList {
		expired := repository.redisClient.TTL(ctx, sessKey).Val()
		sessionTTLList = append(sessionTTLList, sessionTTL{sessKey, expired.Milliseconds()})
	}

	sort.Slice(sessionTTLList, func(i, j int) bool {
		return sessionTTLList[i].exp > sessionTTLList[j].exp
	})

	sessionClearList := make([]string, 0, len(sessionTTLList[repository.cfg.Auth.MaxRefreshSession:]))
	for _, sessOld := range sessionTTLList[repository.cfg.Auth.MaxRefreshSession:] {
		sessionClearList = append(sessionClearList, sessOld.key)
	}

	return repository.redisClient.Del(ctx, sessionClearList...).Val()
}

// Получаем сессию по id
func (repository *sessionRepo) GetSessionByID(ctx context.Context, userID uuid.UUID, sessionID string) (*model.Session, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.GetSessionByID")
	defer span.Finish()

	sessBytes, err := repository.redisClient.Get(ctx, repository.createKey(userID, sessionID)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "sessionRep.GetSessionByID.redisClient.Get")
	}

	sess := &model.Session{}
	if err = json.Unmarshal(sessBytes, &sess); err != nil {
		return nil, errors.Wrap(err, "sessionRepo.GetSessionByID.json.Unmarshal")
	}
	return sess, nil
}

// Обновляем сессию по id
func (repository *sessionRepo) RefreshByID(ctx context.Context, sess *model.Session, newSessionID string, expire int) (string, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.RefreshByID")
	defer span.Finish()

	if err := repository.DeleteByID(ctx, sess.UserID, sess.SessionID); err != nil {
		return "", err
	}

	sess.SessionID = newSessionID
	sessionKey, err := repository.CreateSession(ctx, sess, expire)
	if err != nil {
		return "", err
	}

	return sessionKey, nil
}

// Удаляем сессию по id
func (repository *sessionRepo) DeleteByID(ctx context.Context, userID uuid.UUID, sessionID string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "sessionRepo.DeleteByID")
	defer span.Finish()

	if err := repository.redisClient.Del(ctx, repository.createKey(userID, sessionID)).Err(); err != nil {
		return errors.Wrap(err, "sessionRepo.DeleteByID")
	}
	return nil
}

func (repository *sessionRepo) createKey(userID uuid.UUID, sessionID string) string {
	return fmt.Sprintf("%s%s:%s", repository.basePrefix, userID, sessionID)
}
