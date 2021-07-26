package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/portfolio"
	"github.com/seregaa020292/capitalhub/internal/portfolio/model"
	"github.com/seregaa020292/capitalhub/pkg/logger"
)

// Portfolio UseCase
type portfolioUC struct {
	cfg           *config.Config
	portfolioRepo portfolio.Repository
	log           logger.Logger
}

// Portfolio UseCase constructor
func NewPortfolioUseCase(
	cfg *config.Config,
	portfolioRepo portfolio.Repository,
	log logger.Logger,
) portfolio.UseCase {
	return &portfolioUC{
		cfg:           cfg,
		portfolioRepo: portfolioRepo,
		log:           log,
	}
}

// Создаем новый портфель
func (useCase *portfolioUC) CreateFirst(ctx context.Context, userID uuid.UUID) (*model.Portfolio, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioUC.CreateFirst")
	defer span.Finish()

	currencyId, _ := uuid.Parse(useCase.cfg.Portfolio.CurrencyDefault)

	return useCase.portfolioRepo.Create(ctx, &model.Portfolio{
		Title:      useCase.cfg.Portfolio.TitleDefault,
		Active:     true,
		UserID:     userID,
		CurrencyID: currencyId,
	})
}

func (useCase *portfolioUC) Create(ctx context.Context, portfolio *model.Portfolio) (*model.PortfolioStats, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioUC.Create")
	defer span.Finish()

	portfolioModel, err := useCase.portfolioRepo.Create(ctx, portfolio)
	if err != nil {
		return nil, err
	}

	return useCase.portfolioRepo.GetStats(ctx, portfolioModel.PortfolioID)
}

// Получаем активный портфель по id пользователя
func (useCase *portfolioUC) GetActive(ctx context.Context, userID uuid.UUID) (*model.Portfolio, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioUC.GetActive")
	defer span.Finish()

	return useCase.portfolioRepo.GetActive(ctx, userID)
}

func (useCase *portfolioUC) CheckUserPortfolio(ctx context.Context, userID uuid.UUID, portfolioID uuid.UUID) bool {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioUC.CheckUserPortfolio")
	defer span.Finish()

	if err := useCase.portfolioRepo.CheckUserPortfolio(ctx, userID, portfolioID); err != nil {
		useCase.log.Errorf("portfolioUC.CheckUserPortfolio: %v", err)
		return false
	}

	return true
}

func (useCase *portfolioUC) GetAllStats(ctx context.Context, userID uuid.UUID) (*[]model.PortfolioStats, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioUC.GetAllStats")
	defer span.Finish()

	return useCase.portfolioRepo.GetAllStats(ctx, userID)
}
