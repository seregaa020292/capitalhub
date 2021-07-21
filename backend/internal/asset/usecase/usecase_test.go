package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/require"

	"github.com/seregaa020292/capitalhub/internal/asset/mock"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

func TestAssetUC_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetRepo := mock.NewMockRepository(ctrl)
	assetUC := NewAssetUseCase(nil, mockAssetRepo, apiLogger)

	asset := &models.Asset{}

	span, ctx := opentracing.StartSpanFromContext(context.Background(), "assetUC.Create")
	defer span.Finish()

	mockAssetRepo.EXPECT().Create(ctx, gomock.Eq(asset)).Return(asset, nil)

	createdAsset, err := assetUC.Create(context.Background(), asset)
	require.NoError(t, err)
	require.NotNil(t, createdAsset)
}

func TestAssetUC_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetRepo := mock.NewMockRepository(ctrl)
	assetUC := NewAssetUseCase(nil, mockAssetRepo, apiLogger)

	userUID := uuid.New()

	asset := &models.Asset{
		AssetID: uuid.New(),
		UserID:  userUID,
	}

	baseAsset := &models.AssetBase{
		UserID: userUID,
	}

	user := &models.User{
		UserID: userUID,
	}

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "assetUC.Update")
	defer span.Finish()

	mockAssetRepo.EXPECT().GetByID(ctxWithTrace, gomock.Eq(asset.AssetID)).Return(baseAsset, nil)
	mockAssetRepo.EXPECT().Update(ctxWithTrace, gomock.Eq(asset)).Return(asset, nil)

	updatedAsset, err := assetUC.Update(ctx, asset)
	require.NoError(t, err)
	require.NotNil(t, updatedAsset)
}

func TestAssetUC_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetRepo := mock.NewMockRepository(ctrl)
	assetUC := NewAssetUseCase(nil, mockAssetRepo, apiLogger)

	userUID := uuid.New()

	asset := &models.Asset{
		AssetID: uuid.New(),
		UserID:  userUID,
	}

	baseAsset := &models.AssetBase{
		UserID: userUID,
	}

	user := &models.User{
		UserID: userUID,
	}

	ctx := context.WithValue(context.Background(), utils.UserCtxKey{}, user)
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "assetUC.Delete")
	defer span.Finish()

	mockAssetRepo.EXPECT().GetByID(ctxWithTrace, gomock.Eq(asset.AssetID)).Return(baseAsset, nil)
	mockAssetRepo.EXPECT().Delete(ctxWithTrace, gomock.Eq(asset.AssetID)).Return(nil)

	err := assetUC.Delete(ctx, asset.AssetID)
	require.NoError(t, err)
	require.Nil(t, err)
}

func TestAssetUC_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetRepo := mock.NewMockRepository(ctrl)
	assetUC := NewAssetUseCase(nil, mockAssetRepo, apiLogger)

	asset := &models.Asset{
		AssetID: uuid.New(),
	}

	baseAsset := &models.AssetBase{}

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "assetUC.GetByID")
	defer span.Finish()

	mockAssetRepo.EXPECT().GetByID(ctxWithTrace, gomock.Eq(asset.AssetID)).Return(baseAsset, nil)

	assetBase, err := assetUC.GetByID(ctx, asset.AssetID)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, assetBase)
}

func TestAssetUC_GetAllByMarketID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetRepo := mock.NewMockRepository(ctrl)
	assetUC := NewAssetUseCase(nil, mockAssetRepo, apiLogger)

	marketID := uuid.New()

	asset := &models.Asset{
		AssetID: uuid.New(),
		MarketID:    marketID,
	}

	assetList := &models.AssetList{}

	ctx := context.Background()
	span, ctxWithTrace := opentracing.StartSpanFromContext(ctx, "assetUC.GetAllByMarketID")
	defer span.Finish()

	query := &utils.PaginationQuery{
		Size:    10,
		Page:    1,
		OrderBy: "",
	}

	mockAssetRepo.EXPECT().GetAllByMarketID(ctxWithTrace, gomock.Eq(asset.AssetID), query).Return(assetList, nil)

	assetList, err := assetUC.GetAllByMarketID(ctx, asset.MarketID, query)
	require.NoError(t, err)
	require.Nil(t, err)
	require.NotNil(t, assetList)
}
