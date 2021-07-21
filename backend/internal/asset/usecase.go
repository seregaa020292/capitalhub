//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package asset

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Asset use case
type UseCase interface {
	Create(ctx context.Context, asset *models.Asset) (*models.AssetTotal, error)
	GetAll(ctx context.Context, userID uuid.UUID) (*[]models.AssetBase, error)
	GetTotalAll(ctx context.Context, userID uuid.UUID) (*[]models.AssetTotal, error)
	Update(ctx context.Context, asset *models.Asset) (*models.Asset, error)
	Delete(ctx context.Context, assetID uuid.UUID) error
	GetByID(ctx context.Context, assetID uuid.UUID) (*models.AssetBase, error)
	GetAllByMarketID(ctx context.Context, marketID uuid.UUID, query *utils.PaginationQuery) (*models.AssetList, error)
}
