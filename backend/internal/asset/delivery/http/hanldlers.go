package http

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
	"net/http"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/asset"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Asset handlers
type assetHandlers struct {
	cfg         *config.Config
	assetUC     asset.UseCase
	portfolioUC portfolio.UseCase
	logger      logger.Logger
}

// NewAssetHandlers Asset handlers constructor
func NewAssetHandlers(
	cfg *config.Config,
	assetUC asset.UseCase,
	portfolioUC portfolio.UseCase,
	logger logger.Logger,
) asset.Handlers {
	return &assetHandlers{
		cfg:         cfg,
		assetUC:     assetUC,
		portfolioUC: portfolioUC,
		logger:      logger,
	}
}

// Add
// @Summary Создать новый актив пользователя
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Param input body models.AssetAdd true "Add asset"
// @Success 201 {object} models.Asset
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/add [post]
func (handler *assetHandlers) Add() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.Add")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assetModel := &models.AssetAdd{}
		if err := utils.SanitizeRequest(echoCtx, assetModel); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		handler.portfolioUC.CheckUserPortfolio(ctx, user.UserID, assetModel.PortfolioID)

		createdAsset, err := handler.assetUC.Create(ctx, &models.Asset{
			MarketID:    assetModel.MarketID,
			PortfolioID: assetModel.PortfolioID,
			Amount:      assetModel.Amount,
			Quantity:    assetModel.Quantity,
			NotationAt:  assetModel.NotationAt,
		})
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusCreated, createdAsset)
	}
}

// GetAll
// @Summary Активы пользователя
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Success 201 {object} models.AssetList
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/all [get]
func (handler *assetHandlers) GetAll() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.GetAll")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assets, err := handler.assetUC.GetAll(ctx, user.UserID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, assets)
	}
}

// GetTotalAll
// @Summary Активы пользователя, подсчитанные
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Success 201 {object} []models.AssetTotal
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/total-all [get]
func (handler *assetHandlers) GetTotalAll() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.GetTotalAll")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assets, err := handler.assetUC.GetTotalAll(ctx, user.UserID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, assets)
	}
}

// Update
// @Summary Обновить данные актива
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Param id path int true "asset_id"
// @Success 200 {object} models.Asset
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/{id} [put]
func (handler *assetHandlers) Update() echo.HandlerFunc {
	type UpdateAsset struct {
		Amount   int `json:"amount" db:"amount" validate:"required,gt=0"`
		Quantity int `json:"quantity" db:"quantity" validate:"required,gt=0"`
	}
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.Update")
		defer span.Finish()

		assetID, err := uuid.Parse(echoCtx.Param("asset_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assetModel := &UpdateAsset{}
		if err = utils.SanitizeRequest(echoCtx, assetModel); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		updatedAsset, err := handler.assetUC.Update(ctx, &models.Asset{
			AssetID:  assetID,
			Amount:   assetModel.Amount,
			Quantity: assetModel.Quantity,
		})
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, updatedAsset)
	}
}

// Delete
// @Summary Удалить актив
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Param id path int true "asset_id"
// @Success 200 {string} string	"ok"
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/{id} [delete]
func (handler *assetHandlers) Delete() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.Delete")
		defer span.Finish()

		assetID, err := uuid.Parse(echoCtx.Param("asset_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		if err = handler.assetUC.Delete(ctx, assetID); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary Получить актив по идентификатору
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Param id path int true "asset_id"
// @Success 200 {object} models.Asset
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/{id} [get]
func (handler *assetHandlers) GetByID() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.GetByID")
		defer span.Finish()

		assetID, err := uuid.Parse(echoCtx.Param("asset_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assetModel, err := handler.assetUC.GetByID(ctx, assetID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, assetModel)
	}
}

// GetAllByMarketID
// @Summary Получить все активы по идентификатору ценной бумаги/товара
// @Security Auth
// @Tags Asset
// @Accept json
// @Produce json
// @Param id path int true "market_id"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.AssetList
// @Failure 500 {object} httpErrors.RestErr
// @Router /asset/market/{id} [get]
func (handler *assetHandlers) GetAllByMarketID() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "assetHandlers.GetAllByMarketID")
		defer span.Finish()

		marketID, err := uuid.Parse(echoCtx.Param("market_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		pq, err := utils.GetPaginationFromCtx(echoCtx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assetList, err := handler.assetUC.GetAllByMarketID(ctx, marketID, pq)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, assetList)
	}
}
