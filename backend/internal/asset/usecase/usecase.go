package usecase

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/asset"
	"github.com/seregaa020292/capitalhub/internal/asset/model"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Asset UseCase
type assetUC struct {
	cfg       *config.Config
	assetRepo asset.Repository
	logger    logger.Logger
}

// Asset UseCase constructor
func NewAssetUseCase(cfg *config.Config, assetRepo asset.Repository, logger logger.Logger) asset.UseCase {
	return &assetUC{cfg: cfg, assetRepo: assetRepo, logger: logger}
}

// Create asset
func (u *assetUC) Create(ctx context.Context, asset *model.Asset) (*model.AssetTotal, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.Create")
	defer span.Finish()
	return u.assetRepo.Create(ctx, asset)
}

func (u *assetUC) GetAll(ctx context.Context, userID uuid.UUID) (*[]model.AssetBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.GetAll")
	defer span.Finish()

	return u.assetRepo.GetAll(ctx, userID)
}

func (u *assetUC) GetTotalAll(ctx context.Context, userID uuid.UUID) (*[]model.AssetTotal, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.GetTotalAll")
	defer span.Finish()

	return u.assetRepo.GetTotalAll(ctx, userID)
}

// Update asset
func (u *assetUC) Update(ctx context.Context, asset *model.Asset) (*model.Asset, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.Update")
	defer span.Finish()

	assetModel, err := u.assetRepo.GetByID(ctx, asset.AssetID)
	if err != nil {
		return nil, err
	}

	if err = utils.ValidateIsOwner(ctx, assetModel.UserID.String(), u.logger); err != nil {
		return nil, httpErrors.NewRestError(http.StatusForbidden, "Forbidden", errors.Wrap(err, "assetUC.Update.ValidateIsOwner"))
	}

	updatedAsset, err := u.assetRepo.Update(ctx, asset)
	if err != nil {
		return nil, err
	}

	return updatedAsset, nil
}

// Delete asset
func (u *assetUC) Delete(ctx context.Context, assetID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.Delete")
	defer span.Finish()

	assetModel, err := u.assetRepo.GetByID(ctx, assetID)
	if err != nil {
		return err
	}

	if err = utils.ValidateIsOwner(ctx, assetModel.UserID.String(), u.logger); err != nil {
		return httpErrors.NewRestError(http.StatusForbidden, "Forbidden", errors.Wrap(err, "assetUC.Delete.ValidateIsOwner"))
	}

	if err = u.assetRepo.Delete(ctx, assetID); err != nil {
		return err
	}

	return nil
}

// GetByID asset
func (u *assetUC) GetByID(ctx context.Context, assetID uuid.UUID) (*model.AssetBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.GetByID")
	defer span.Finish()

	return u.assetRepo.GetByID(ctx, assetID)
}

// GetAllByMarketID asset
func (u *assetUC) GetAllByMarketID(ctx context.Context, marketID uuid.UUID, query *utils.PaginationQuery) (*model.AssetList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetUC.GetAllByMarketID")
	defer span.Finish()

	return u.assetRepo.GetAllByMarketID(ctx, marketID, query)
}
