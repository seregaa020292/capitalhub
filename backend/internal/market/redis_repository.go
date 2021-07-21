//go:generate mockgen -source redis_repository.go -destination mock/redis_repository_mock.go -package mock
package market

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/market/model"
)

// Market redis repository
type RedisRepository interface {
	GetMarketByIDCtx(ctx context.Context, key string) (*model.MarketBase, error)
	SetMarketCtx(ctx context.Context, key string, seconds int, asset *model.MarketBase) error
	DeleteMarketCtx(ctx context.Context, key string) error
}
