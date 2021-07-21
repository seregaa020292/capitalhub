package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/market"
	"github.com/seregaa020292/capitalhub/internal/market/service"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Market handlers
type marketHandlers struct {
	cfg          *config.Config
	marketUC     market.UseCase
	parseService *service.ParseService
	logger       logger.Logger
}

// NewMarketHandlers Market handlers constructor
func NewMarketHandlers(
	cfg *config.Config,
	marketUC market.UseCase,
	parseService *service.ParseService,
	logger logger.Logger,
) market.Handlers {
	return &marketHandlers{
		cfg:          cfg,
		marketUC:     marketUC,
		parseService: parseService,
		logger:       logger,
	}
}

// Create godoc
// @Summary Добавить новую ценную бумаги/товара
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Success 201 {object} models.Market
// @Router /market/create [post]
func (handler marketHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "marketHandlers.Create")
		defer span.Finish()

		marketModel := &models.Market{}
		if err := c.Bind(marketModel); err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		createdMarkets, err := handler.marketUC.Create(ctx, marketModel)
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		return c.JSON(http.StatusCreated, createdMarkets)
	}
}

// Update godoc
// @Summary Изменить данные ценной бумаги/товара
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Param id path int true "market_id"
// @Success 200 {object} models.Market
// @Router /market/{id} [put]
func (handler marketHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "marketHandlers.Update")
		defer span.Finish()

		marketUUID, err := uuid.Parse(c.Param("market_id"))
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		marketModel := &models.Market{}
		if err = c.Bind(marketModel); err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}
		marketModel.MarketID = marketUUID

		updatedMarket, err := handler.marketUC.Update(ctx, marketModel)
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		return c.JSON(http.StatusOK, updatedMarket)
	}
}

// GetByID godoc
// @Summary Получить по id ценную бумагу/товар
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Param id path int true "market_id"
// @Success 200 {object} models.Market
// @Router /market/{id} [get]
func (handler marketHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "marketHandlers.GetByID")
		defer span.Finish()

		marketUUID, err := uuid.Parse(c.Param("market_id"))
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		marketByID, err := handler.marketUC.GetByID(ctx, marketUUID)
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		return c.JSON(http.StatusOK, marketByID)
	}
}

// Delete godoc
// @Summary Удалить ценную бумагу/товар
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Param id path int true "market_id"
// @Success 200 {string} string	"ok"
// @Router /market/{id} [delete]
func (handler marketHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "marketHandlers.Delete")
		defer span.Finish()

		marketUUID, err := uuid.Parse(c.Param("market_id"))
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		if err = handler.marketUC.Delete(ctx, marketUUID); err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetMarkets godoc
// @Summary Получить все ценные бумаги/товары
// @Description Получить все ценные бумаги/товары с разбивкой на страницы
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.MarketList
// @Router /market/all [get]
func (handler marketHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "marketHandlers.GetAll")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		marketsList, err := handler.marketUC.GetAll(ctx, pq)
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		return c.JSON(http.StatusOK, marketsList)
	}
}

// GetMarkets godoc
// @Summary Спарсить
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Success 200
// @Router /market/parse [get]
func (handler marketHandlers) Parse() echo.HandlerFunc {
	return func(ctxEcho echo.Context) error {
		span, _ := opentracing.StartSpanFromContext(utils.GetRequestCtx(ctxEcho), "marketHandlers.Parse")
		defer span.Finish()

		if err := handler.parseService.TCSParse("currencies"); err != nil {
			return utils.ErrResponseWithLog(ctxEcho, handler.logger, err)
		}

		return ctxEcho.NoContent(http.StatusOK)
	}
}

// SearchByTitle godoc
// @Summary Поиск по названию
// @Description Искать ценную бумагу/товар по названию
// @Security Auth
// @Tags Market
// @Accept json
// @Produce json
// @Param title query int false "title"
// @Param page query int false "page number" Format(page)
// @Param size query int false "number of elements per page" Format(size)
// @Param orderBy query int false "filter name" Format(orderBy)
// @Success 200 {object} models.MarketList
// @Router /market/search [get]
func (handler marketHandlers) SearchByTitle() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "marketHandlers.SearchByTitle")
		defer span.Finish()

		pq, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		marketsList, err := handler.marketUC.SearchByTitle(ctx, c.QueryParam("title"), pq)

		if err != nil {
			return utils.ErrResponseWithLog(c, handler.logger, err)
		}

		return c.JSON(http.StatusOK, marketsList)
	}
}
