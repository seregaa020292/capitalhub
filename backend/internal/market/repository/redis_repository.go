package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/market"
	"github.com/seregaa020292/capitalhub/internal/models"
)

// Market redis repository
type marketRedisRepo struct {
	redisClient *redis.Client
}

// Market redis repository constructor
func NewMarketRedisRepository(redisClient *redis.Client) market.RedisRepository {
	return &marketRedisRepo{redisClient: redisClient}
}

// Get market by id
func (n *marketRedisRepo) GetMarketByIDCtx(ctx context.Context, key string) (*models.MarketBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRedisRepo.GetMarketByIDCtx")
	defer span.Finish()

	marketBytes, err := n.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "marketRedisRepo.GetMarketByIDCtx.redisClient.Get")
	}
	marketBase := &models.MarketBase{}
	if err = json.Unmarshal(marketBytes, marketBase); err != nil {
		return nil, errors.Wrap(err, "marketRedisRepo.GetMarketByIDCtx.json.Unmarshal")
	}

	return marketBase, nil
}

// Cache market item
func (n *marketRedisRepo) SetMarketCtx(ctx context.Context, key string, seconds int, market *models.MarketBase) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRedisRepo.SetMarketCtx")
	defer span.Finish()

	marketBytes, err := json.Marshal(market)
	if err != nil {
		return errors.Wrap(err, "marketRedisRepo.SetMarketCtx.json.Marshal")
	}
	if err = n.redisClient.Set(ctx, key, marketBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "marketRedisRepo.SetMarketCtx.redisClient.Set")
	}
	return nil
}

// Delete market item from cache
func (n *marketRedisRepo) DeleteMarketCtx(ctx context.Context, key string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketRedisRepo.DeleteMarketCtx")
	defer span.Finish()

	if err := n.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "marketRedisRepo.DeleteMarketCtx.redisClient.Del")
	}
	return nil
}
