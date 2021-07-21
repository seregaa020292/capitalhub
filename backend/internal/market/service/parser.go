package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/currency"
	"github.com/seregaa020292/capitalhub/internal/instrument"
	"github.com/seregaa020292/capitalhub/internal/market"
	"github.com/seregaa020292/capitalhub/internal/market/model"
	"github.com/seregaa020292/capitalhub/internal/provider"
	"github.com/seregaa020292/capitalhub/internal/register"
	registerModel "github.com/seregaa020292/capitalhub/internal/register/model"
)

type ParseService struct {
	cfg          *config.Config
	marketUC     market.UseCase
	providerUC   provider.UseCase
	instrumentUC instrument.UseCase
	currencyUC   currency.UseCase
	registerUC   register.UseCase
}

func NewTCSParseService(
	cfg *config.Config,
	marketUC market.UseCase,
	providerUC provider.UseCase,
	instrumentUC instrument.UseCase,
	currencyUC currency.UseCase,
	registerUC register.UseCase,
) *ParseService {
	return &ParseService{
		cfg:          cfg,
		marketUC:     marketUC,
		providerUC:   providerUC,
		instrumentUC: instrumentUC,
		currencyUC:   currencyUC,
		registerUC:   registerUC,
	}
}

type stock struct {
	Figi              string  `json:"figi"`
	Ticker            string  `json:"ticker"`
	Isin              string  `json:"isin"`
	MinPriceIncrement float32 `json:"minPriceIncrement"`
	Lot               int     `json:"lot"`
	Currency          string  `json:"currency"`
	Name              string  `json:"name"`
	Type              string  `json:"type"`
}

type respStock struct {
	TrackingId string `json:"trackingId"`
	Status     string `json:"status"`
	Payload    struct {
		Total       int     `json:"total"`
		Instruments []stock `json:"instruments"`
	} `json:"payload"`
}

// https://tinkoffcreditsystems.github.io/invest-openapi/swagger-ui/#/
// stocks, bonds, etfs, currencies
func (service ParseService) TCSParse(entity string) error {
	ctx := context.Background()

	var respStock respStock
	if err := service.httpRequest(&respStock, entity); err != nil {
		return err
	}

	type mapUUID map[string]uuid.UUID
	marketTypes, err := service.instrumentUC.GetAll(ctx)
	if err != nil {
		return err
	}
	marketTypeMap := make(mapUUID)
	for _, r := range *marketTypes {
		marketTypeMap[r.Title] = r.InstrumentID
	}

	currencies, err := service.currencyUC.GetAll(ctx)
	if err != nil {
		return err
	}
	currencyMap := make(mapUUID)
	for _, r := range *currencies {
		currencyMap[r.Title] = r.CurrencyID
	}

	providerModel, err := service.providerUC.GetByTitle(ctx, "Tinkoff")
	if err != nil {
		return err
	}

	for _, stock := range respStock.Payload.Instruments {
		marketModel, err := service.marketUC.Create(ctx, &model.Market{
			Title:        stock.Name,
			Ticker:       stock.Ticker,
			Content:      "",
			ImageURL:     nil,
			CurrencyID:   currencyMap[stock.Currency],
			InstrumentID: marketTypeMap[strings.ToUpper(stock.Type)],
		})
		if err != nil {
			if strings.Contains(err.Error(), "23505") {
				fmt.Println("DUPLICATE:", stock)
				continue
			} else {
				return err
			}
		}

		if _, err := service.registerUC.Create(ctx, &registerModel.Register{
			Identify:   stock.Figi,
			ProviderID: providerModel.ProviderID,
			MarketID:   marketModel.MarketID,
		}); err != nil {
			if strings.Contains(err.Error(), "23505") {
				fmt.Println("DUPLICATE:", stock)
			} else {
				return err
			}
		}
	}

	return nil
}

func (service ParseService) httpRequest(respStock *respStock, entity string) error {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("https://api-invest.tinkoff.ru/openapi/sandbox/market/%s", entity),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+service.cfg.TCS.Token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(resBody, &respStock)
}
