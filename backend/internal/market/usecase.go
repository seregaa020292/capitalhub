//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package market

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Market use case
type UseCase interface {
	Create(ctx context.Context, market *models.Market) (*models.Market, error)
	Update(ctx context.Context, market *models.Market) (*models.Market, error)
	GetByID(ctx context.Context, marketID uuid.UUID) (*models.MarketBase, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*[]models.MarketRegister, error)
	Delete(ctx context.Context, marketID uuid.UUID) error
	GetAll(ctx context.Context, pq *utils.PaginationQuery) (*models.MarketList, error)
	SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.MarketList, error)
}
