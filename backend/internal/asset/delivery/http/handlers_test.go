package http

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"github.com/seregaa020292/capitalhub/internal/asset/mock"
	"github.com/seregaa020292/capitalhub/internal/asset/model"
	"github.com/seregaa020292/capitalhub/internal/asset/usecase"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/converter"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

func TestAssetHandlers_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetUC := mock.NewMockUseCase(ctrl)
	assetUC := usecase.NewAssetUseCase(nil, mockAssetUC, apiLogger)

	assetHandlers := NewAssetHandlers(nil, assetUC, apiLogger)
	handlerFunc := assetHandlers.Create()

	userID := uuid.New()
	assetUID := uuid.New()
	asset := &model.Asset{
		UserID:  userID,
		Message: "message Key: 'Asset.Message' Error:Field validation for 'Message' failed on the 'gte' tag",
		AssetID: assetUID,
	}

	buf, err := converter.AnyToBytesBuffer(asset)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/assets", strings.NewReader(buf.String()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	u := &models.User{
		UserID: userID,
	}
	ctxWithValue := context.WithValue(context.Background(), utils.UserCtxKey{}, u)
	req = req.WithContext(ctxWithValue)

	e := echo.New()
	ctx := e.NewContext(req, res)

	mockAsset := &model.Asset{
		UserID:  userID,
		AssetID: asset.AssetID,
		Message: "message",
	}

	fmt.Printf("ASSET: %#v\n", asset)
	fmt.Printf("MOCK ASSET: %#v\n", mockAsset)

	mockAssetUC.EXPECT().Create(gomock.Any(), gomock.Any()).Return(mockAsset, nil)

	err = handlerFunc(ctx)
	require.NoError(t, err)
}

func TestAssetHandlers_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetUC := mock.NewMockUseCase(ctrl)
	assetUC := usecase.NewAssetUseCase(nil, mockAssetUC, apiLogger)

	assetHandlers := NewAssetHandlers(nil, assetUC, apiLogger)
	handlerFunc := assetHandlers.GetByID()

	r := httptest.NewRequest(http.MethodGet, "/api/v1/assets/5c9a9d67-ad38-499c-9858-086bfdeaf7d2", nil)
	w := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(r, w)
	c.SetParamNames("asset_id")
	c.SetParamValues("5c9a9d67-ad38-499c-9858-086bfdeaf7d2")

	asset := &model.AssetBase{}

	mockAssetUC.EXPECT().GetByID(gomock.Any(), gomock.Any()).Return(asset, nil)

	err := handlerFunc(c)
	require.NoError(t, err)
}

func TestAssetHandlers_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	apiLogger := logger.NewApiLogger(nil)
	mockAssetUC := mock.NewMockUseCase(ctrl)
	assetUC := usecase.NewAssetUseCase(nil, mockAssetUC, apiLogger)

	assetHandlers := NewAssetHandlers(nil, assetUC, apiLogger)
	handlerFunc := assetHandlers.Delete()

	userID := uuid.New()
	assetID := uuid.New()
	asset := &model.AssetBase{
		AssetID: assetID,
		UserID:  userID,
	}

	r := httptest.NewRequest(http.MethodDelete, "/api/v1/assets/5c9a9d67-ad38-499c-9858-086bfdeaf7d2", nil)
	w := httptest.NewRecorder()
	u := &models.User{
		UserID: userID,
	}
	ctxWithValue := context.WithValue(context.Background(), utils.UserCtxKey{}, u)
	r = r.WithContext(ctxWithValue)
	e := echo.New()
	c := e.NewContext(r, w)
	c.SetParamNames("asset_id")
	c.SetParamValues(assetID.String())

	mockAssetUC.EXPECT().GetByID(gomock.Any(), assetID).Return(asset, nil)
	mockAssetUC.EXPECT().Delete(gomock.Any(), assetID).Return(nil)

	err := handlerFunc(c)
	require.NoError(t, err)
}
