package usecase

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/register"
	"github.com/seregaa020292/capitalhub/internal/register/model"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

type registerUC struct {
	cfg          *config.Config
	registerRepo register.Repository
	logger       logger.Logger
}

// Register UseCase constructor
func NewRegisterUseCase(
	cfg *config.Config,
	registerRepo register.Repository,
	logger logger.Logger,
) register.UseCase {
	return &registerUC{
		cfg:          cfg,
		registerRepo: registerRepo,
		logger:       logger,
	}
}

func (useCase *registerUC) Create(ctx context.Context, register *model.Register) (*model.Register, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "registerUC.Create")
	defer span.Finish()

	if err := utils.ValidateStruct(ctx, register); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "registerUC.Create.ValidateStruct"))
	}

	n, err := useCase.registerRepo.Create(ctx, register)
	if err != nil {
		return nil, err
	}

	return n, err
}
