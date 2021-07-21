package socket

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/market"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/quote"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

type marketHandlers struct {
	cfg          *config.Config
	marketUC     market.UseCase
	quotesClient quote.Listener
	logger       logger.Logger
}

// NewMarketHandlers Market handlers constructor
func NewMarketHandlers(
	cfg *config.Config,
	marketUC market.UseCase,
	quotesClient quote.Listener,
	logger logger.Logger,
) market.SocketHandlers {
	return &marketHandlers{
		cfg:          cfg,
		marketUC:     marketUC,
		quotesClient: quotesClient,
		logger:       logger,
	}
}

func (handler marketHandlers) Quotes() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(echoCtx), "marketHandlers.Quotes")
		defer span.Finish()

		user, err := utils.GetUserFromCtx(ctx)
		if err != nil {
			return utils.ErrResponseWithLog(echoCtx, handler.logger, err)
		}

		marketRegisters, err := handler.marketUC.GetByUserID(ctx, user.UserID)
		if err != nil {
			handler.logger.Error(err)
			return nil
		}

		return handler.quotesClient.Subscribe(echoCtx, *marketRegisters)
	}
}
