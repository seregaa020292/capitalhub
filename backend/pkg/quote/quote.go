package quote

import (
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"

	"github.com/labstack/echo/v4"
	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/market/model"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
	"github.com/seregaa020292/capitalhub/pkg/webSocket"
)

type Listener interface {
	Subscribe(echo.Context, []model.MarketRegister) error
	Terminate()
}

type quote struct {
	webSocketClient *webSocket.Hub
	tcsClient       *sdk.StreamingClient
	cfg             *config.Config
	logger          logger.Logger
}

func NewClient(
	webSocketClient *webSocket.Hub,
	cfg *config.Config,
	logger logger.Logger,
) Listener {
	return quote{
		webSocketClient: webSocketClient,
		// Setup Tinkoff Client
		tcsClient: NewTCSClient(webSocketClient, cfg.TCS, logger),
		cfg:       cfg,
		logger:    logger,
	}
}

func (q quote) Subscribe(echoCtx echo.Context, marketRegisters []model.MarketRegister) error {
	conn, err := q.webSocketClient.NewConnection(echoCtx.Response(), echoCtx.Request())
	if err != nil {
		q.logger.Error(err)
		return err
	}

	for _, register := range marketRegisters {
		q.logger.Infof("Подписка на получение свечей по инструменту %s (%s)", register.Identify, register.Ticker)
		if err = q.tcsClient.SubscribeCandle(register.Identify, sdk.CandleInterval5Min, utils.RequestID()); err != nil {
			q.logger.Error(err)
		}
	}

	select {
	case message := <-conn.Messages():
		q.logger.Infof("recv: %s", message)
	}

	return nil
}

func (q quote) Terminate() {
	q.tcsClient.Close()
}
