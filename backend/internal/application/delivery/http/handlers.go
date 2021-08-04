package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/application"
	"github.com/seregaa020292/capitalhub/internal/application/model"
	"github.com/seregaa020292/capitalhub/internal/currency"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

type applicationHandlers struct {
	cfg           *config.Config
	applicationUC application.UseCase
	currencyUC    currency.UseCase
	logger        logger.Logger
}

func NewApplicationHandlers(
	cfg *config.Config,
	applicationUC application.UseCase,
	currencyUC currency.UseCase,
	logger logger.Logger,
) application.Handlers {
	return &applicationHandlers{
		cfg:           cfg,
		applicationUC: applicationUC,
		currencyUC:    currencyUC,
		logger:        logger,
	}
}

// GetDashboard
// @Summary Общие данные для панели
// @Security Auth
// @Tags Application
// @Accept json
// @Produce json
// @Success 200 {object} model.Dashboard
// @Router /application/dashboard [get]
func (handler *applicationHandlers) GetDashboard() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "applicationHandlers.GetDashboard")
		defer span.Finish()

		currencies, err := handler.currencyUC.GetAll(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		return echoCtx.JSON(http.StatusOK, model.Dashboard{
			Currencies: currencies,
		})
	}
}
