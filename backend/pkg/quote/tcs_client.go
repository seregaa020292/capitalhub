package quote

import (
	"encoding/json"
	"log"
	"os"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/webSocket"
)

func NewTCSClient(webSocketClient *webSocket.Hub, tcsCfg config.TCSConfig, logger logger.Logger) *sdk.StreamingClient {
	streamLogger := log.New(os.Stdout, "[invest-openapi-go-sdk]", log.LstdFlags)

	client, err := sdk.NewStreamingClient(streamLogger, tcsCfg.Token)
	if err != nil {
		logger.Error(err)
	}

	// Запускаем цикл обработки входящих событий. Запускаем асинхронно
	// Сюда будут приходить сообщения по подпискам после вызова соответствующих методов
	// SubscribeInstrumentInfo, SubscribeCandle, SubscribeOrderbook
	go func() {
		err = client.RunReadLoop(func(event interface{}) error {
			switch sdkEvent := event.(type) {
			case sdk.CandleEvent:
				msg, err := json.Marshal(sdkEvent)
				if err != nil {
					return err
				}
				webSocketClient.SendMessage(msg)
			}

			return nil
		})

		logger.Info("TCS disconnected")
		if err != nil {
			logger.Error(err)
		}
	}()

	logger.Info("TCS connected")

	return client
}
