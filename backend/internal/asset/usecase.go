//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package asset

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/asset/model"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Asset use case
type UseCase interface {
	Create(ctx context.Context, asset *model.Asset) (*model.AssetTotal, error)
	GetAll(ctx context.Context, userID uuid.UUID) (*[]model.AssetBase, error)
	GetTotalAll(ctx context.Context, userID uuid.UUID) (*[]model.AssetTotal, error)
	Update(ctx context.Context, asset *model.Asset) (*model.Asset, error)
	Delete(ctx context.Context, assetID uuid.UUID) error
	GetByID(ctx context.Context, assetID uuid.UUID) (*model.AssetBase, error)
	GetAllByMarketID(ctx context.Context, marketID uuid.UUID, query *utils.PaginationQuery) (*model.AssetList, error)
}
