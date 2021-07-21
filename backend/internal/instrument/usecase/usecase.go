package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/instrument"
	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/pkg/logger"
)

// Instrument UseCase
type instrumentUC struct {
	cfg            *config.Config
	instrumentRepo instrument.Repository
	logger         logger.Logger
}

// Instrument UseCase constructor
func NewInstrumentUseCase(
	cfg *config.Config,
	instrumentRepo instrument.Repository,
	logger logger.Logger,
) instrument.UseCase {
	return &instrumentUC{
		cfg:            cfg,
		instrumentRepo: instrumentRepo,
		logger:         logger,
	}
}

// Get all instruments
func (useCase *instrumentUC) GetAll(ctx context.Context) (*[]models.Instrument, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "instrumentUC.GetAll")
	defer span.Finish()

	return useCase.instrumentRepo.GetAll(ctx)
}
