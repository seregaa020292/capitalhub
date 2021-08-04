package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/application"
	"github.com/seregaa020292/capitalhub/internal/application/model"
	"github.com/seregaa020292/capitalhub/pkg/logger"
)

type applicationUC struct {
	cfg             *config.Config
	applicationRepo application.Repository
	logger          logger.Logger
}

func NewRegisterUseCase(
	cfg *config.Config,
	applicationRepo application.Repository,
	logger logger.Logger,
) application.UseCase {
	return &applicationUC{
		cfg:             cfg,
		applicationRepo: applicationRepo,
		logger:          logger,
	}
}

func (useCase applicationUC) GetDashboard(ctx context.Context) (*model.Dashboard, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "applicationUC.GetDashboard")
	defer span.Finish()

	return useCase.applicationRepo.GetDashboard(ctx)
}
