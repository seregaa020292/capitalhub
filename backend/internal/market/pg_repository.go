//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package market

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/market/model"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Market Repository
type Repository interface {
	Create(ctx context.Context, market *model.Market) (*model.Market, error)
	Update(ctx context.Context, market *model.Market) (*model.Market, error)
	GetByID(ctx context.Context, marketID uuid.UUID) (*model.MarketBase, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*[]model.MarketRegister, error)
	Delete(ctx context.Context, marketID uuid.UUID) error
	GetAll(ctx context.Context, pq *utils.PaginationQuery) (*model.MarketList, error)
	SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*model.MarketList, error)
}
