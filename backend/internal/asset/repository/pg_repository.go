package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/asset"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Asset Repository
type assetRepo struct {
	db *sqlx.DB
}

// Asset Repository constructor
func NewAssetRepository(db *sqlx.DB) asset.Repository {
	return &assetRepo{db: db}
}

// Create asset
func (repository *assetRepo) Create(ctx context.Context, asset *models.Asset) (*models.AssetTotal, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.Create")
	defer span.Finish()

	assetModel := &models.Asset{}
	if err := repository.db.QueryRowxContext(
		ctx,
		createAsset,
		&asset.PortfolioID,
		&asset.MarketID,
		&asset.Amount,
		&asset.Quantity,
		&asset.NotationAt,
	).StructScan(assetModel); err != nil {
		return nil, errors.Wrap(err, "assetRepo.Create.StructScan")
	}

	assetTotalModel := &models.AssetTotal{}
	if err := repository.db.GetContext(ctx, assetTotalModel, getTotalAssetByMarketID, assetModel.MarketID, asset.PortfolioID); err != nil {
		return nil, errors.Wrap(err, "assetRepo.Create.GetContext")
	}

	return assetTotalModel, nil
}

// GetAll asset
func (repository *assetRepo) GetAll(ctx context.Context, userID uuid.UUID) (*[]models.AssetBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.GetAll")
	defer span.Finish()

	assets := &[]models.AssetBase{}
	if err := repository.db.SelectContext(ctx, assets, getAssetByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "assetRepo.GetAll.GetContext")
	}
	return assets, nil
}

// GetTotalAll asset
func (repository *assetRepo) GetTotalAll(ctx context.Context, userID uuid.UUID) (*[]models.AssetTotal, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.GetTotalAll")
	defer span.Finish()

	assets := &[]models.AssetTotal{}
	if err := repository.db.SelectContext(ctx, assets, getTotalAssetByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "assetRepo.GetTotalAll.GetContext")
	}
	return assets, nil
}

// Update asset
func (repository *assetRepo) Update(ctx context.Context, asset *models.Asset) (*models.Asset, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.Update")
	defer span.Finish()

	assetModel := &models.Asset{}
	if err := repository.db.QueryRowxContext(ctx, updateAsset, asset.Amount, asset.Quantity, asset.AssetID).StructScan(assetModel); err != nil {
		return nil, errors.Wrap(err, "assetRepo.Update.QueryRowxContext")
	}

	return assetModel, nil
}

// Delete asset
func (repository *assetRepo) Delete(ctx context.Context, assetID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.Delete")
	defer span.Finish()

	result, err := repository.db.ExecContext(ctx, deleteAsset, assetID)
	if err != nil {
		return errors.Wrap(err, "assetRepo.Delete.ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "assetRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "assetRepo.Delete.rowsAffected")
	}

	return nil
}

// GetByID asset
func (repository *assetRepo) GetByID(ctx context.Context, assetID uuid.UUID) (*models.AssetBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.GetByID")
	defer span.Finish()

	assetModel := &models.AssetBase{}
	if err := repository.db.GetContext(ctx, assetModel, getAssetByID, assetID); err != nil {
		return nil, errors.Wrap(err, "assetRepo.GetByID.GetContext")
	}
	return assetModel, nil
}

// GetAllByMarketID Asset
func (repository *assetRepo) GetAllByMarketID(ctx context.Context, assetID uuid.UUID, query *utils.PaginationQuery) (*models.AssetList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "assetRepo.GetAllByMarketID")
	defer span.Finish()

	var totalCount int
	if err := repository.db.QueryRowContext(ctx, getTotalCountByAssetID, assetID).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "assetRepo.GetAllByMarketID.QueryRowContext")
	}
	if totalCount == 0 {
		return &models.AssetList{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
			Page:       query.GetPage(),
			Size:       query.GetSize(),
			HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
			Assets:     make([]*models.AssetBase, 0),
		}, nil
	}

	rows, err := repository.db.QueryxContext(ctx, getAssetsByAssetID, assetID, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "assetRepo.GetAllByMarketID.QueryxContext")
	}
	defer rows.Close()

	AssetList := make([]*models.AssetBase, 0, query.GetSize())
	for rows.Next() {
		assetModel := &models.AssetBase{}
		if err = rows.StructScan(assetModel); err != nil {
			return nil, errors.Wrap(err, "assetRepo.GetAllByMarketID.StructScan")
		}
		AssetList = append(AssetList, assetModel)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "assetRepo.GetAllByMarketID.rows.Err")
	}

	return &models.AssetList{
		TotalCount: totalCount,
		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
		Page:       query.GetPage(),
		Size:       query.GetSize(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
		Assets:      AssetList,
	}, nil
}
