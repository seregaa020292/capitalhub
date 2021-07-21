package http

import (
	"github.com/seregaa020292/capitalhub/internal/portfolio/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/asset"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
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
