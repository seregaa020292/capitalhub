package http

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/asset"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
	"github.com/seregaa020292/capitalhub/internal/portfolio/model"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

// Portfolio handlers
type portfolioHandlers struct {
	cfg         *config.Config
	assetUC     asset.UseCase
	portfolioUC portfolio.UseCase
	logger      logger.Logger
}

// NewPortfolioHandlers Portfolio handlers constructor
func NewPortfolioHandlers(
	cfg *config.Config,
	assetUC asset.UseCase,
	portfolioUC portfolio.UseCase,
	logger logger.Logger,
) portfolio.Handlers {
	return &portfolioHandlers{
		cfg:         cfg,
		assetUC:     assetUC,
		portfolioUC: portfolioUC,
		logger:      logger,
	}
}

// GetActiveTotal
// @Summary Портфель пользователя
// @Security Auth
// @Tags Portfolio
// @Accept json
// @Produce json
// @Success 200 {object} model.PortfolioList
// @Failure 500 {object} httpErrors.RestErr
// @Router /portfolio/active-total [get]
func (handler *portfolioHandlers) GetActiveTotal() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "portfolioHandlers.GetActiveTotal")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		portfolioActive, err := handler.portfolioUC.GetActive(ctx, user.UserID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		assets, err := handler.assetUC.GetTotalAll(ctx, user.UserID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, model.PortfolioList{
			Portfolio:  *portfolioActive,
			AssetTotal: *assets,
		})
	}
}

// GetAllStats
// @Summary Портфели пользователя
// @Security Auth
// @Tags Portfolio
// @Accept json
// @Produce json
// @Success 200 {object} []model.PortfolioStats
// @Router /portfolio/all-stats [get]
func (handler *portfolioHandlers) GetAllStats() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "portfolioHandlers.GetAllStats")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		portfolios, err := handler.portfolioUC.GetAllStats(ctx, user.UserID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, portfolios)
	}
}

// Add
// @Summary Портфели пользователя
// @Security Auth
// @Tags Portfolio
// @Accept json
// @Produce json
// @Param input body model.PortfolioChange true "Add portfolio"
// @Success 200 {object} model.PortfolioStats
// @Router /portfolio/add [get]
func (handler *portfolioHandlers) Add() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "portfolioHandlers.Add")
		defer span.Finish()

		portfolioChange := &model.PortfolioChange{}
		if err := utils.ReadRequest(echoCtx, portfolioChange); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		portfolioStats, err := handler.portfolioUC.Create(ctx, &model.Portfolio{
			UserID:     user.UserID,
			Title:      portfolioChange.Title,
			CurrencyID: portfolioChange.CurrencyID,
		})
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, portfolioStats)
	}
}

// Choose
// @Summary Сменить активный портфель пользователя
// @Security Auth
// @Tags Portfolio
// @Accept json
// @Produce json
// @Success 200 {object} bool
// @Router /portfolio/{portfolio_id}/choose [put]
func (handler *portfolioHandlers) Choose() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "portfolioHandlers.Choose")
		defer span.Finish()

		portfolioID, err := uuid.Parse(echoCtx.Param("portfolio_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		choosePortfolio, err := handler.portfolioUC.Choose(ctx, portfolioID, user.UserID)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, choosePortfolio)
	}
}

// Edit
// @Summary Изменение данных портфеля
// @Security Auth
// @Tags Portfolio
// @Accept json
// @Produce json
// @Param input body model.PortfolioChange true "Edit portfolio"
// @Success 200 {object} model.PortfolioStats
// @Router /portfolio/{portfolio_id} [put]
func (handler *portfolioHandlers) Edit() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "portfolioHandlers.Edit")
		defer span.Finish()

		portfolioID, err := uuid.Parse(echoCtx.Param("portfolio_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		portfolioChange := &model.PortfolioChange{}
		if err := utils.ReadRequest(echoCtx, portfolioChange); err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		portfolioEdit, err := handler.portfolioUC.Edit(ctx, portfolioID, user.UserID, portfolioChange)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, portfolioEdit)
	}
}

// Remove
// @Summary Удаление портфеля
// @Security Auth
// @Tags Portfolio
// @Accept json
// @Produce json
// @Success 200 {object} bool
// @Router /portfolio/{portfolio_id} [delete]
func (handler *portfolioHandlers) Remove() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "portfolioHandlers.Remove")
		defer span.Finish()

		portfolioID, err := uuid.Parse(echoCtx.Param("portfolio_id"))
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		portfolioRemoved := handler.portfolioUC.Remove(ctx, portfolioID, user.UserID)

		return echoCtx.JSON(http.StatusOK, portfolioRemoved)
	}
}
