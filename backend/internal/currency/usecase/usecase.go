package usecase

import (
	"context"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/currency"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/logger"
)

// Currency UseCase
type currencyUC struct {
	cfg          *config.Config
	currencyRepo currency.Repository
	logger       logger.Logger
}

// Currency UseCase constructor
func NewCurrencyUseCase(
	cfg *config.Config,
	currencyRepo currency.Repository,
	logger logger.Logger,
) currency.UseCase {
	return &currencyUC{
		cfg:          cfg,
		currencyRepo: currencyRepo,
		logger:       logger,
	}
}

// Get currencies
func (useCase *currencyUC) GetAll(ctx context.Context) (*[]models.Currency, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "currencyUC.GetAll")
	defer span.Finish()

	return useCase.currencyRepo.GetAll(ctx)
}
