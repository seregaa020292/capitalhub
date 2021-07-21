package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/provider"
	"github.com/seregaa020292/capitalhub/internal/provider/model"
	"github.com/seregaa020292/capitalhub/pkg/logger"
)

// Provider UseCase
type providerUC struct {
	cfg          *config.Config
	providerRepo provider.Repository
	logger       logger.Logger
}

// Provider UseCase constructor
func NewProviderUseCase(
	cfg *config.Config,
	providerRepo provider.Repository,
	logger logger.Logger,
) provider.UseCase {
	return &providerUC{
		cfg:          cfg,
		providerRepo: providerRepo,
		logger:       logger,
	}
}

// GetByTitle provider
func (u *providerUC) GetByTitle(ctx context.Context, title string) (*model.Provider, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "providerUC.GetByTitle")
	defer span.Finish()

	return u.providerRepo.GetByTitle(ctx, title)
}
