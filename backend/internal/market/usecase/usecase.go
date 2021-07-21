package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/config"
	"github.com/seregaa020292/capitalhub/internal/market"
	"github.com/seregaa020292/capitalhub/internal/market/model"
	"github.com/seregaa020292/capitalhub/pkg/httpErrors"
	"github.com/seregaa020292/capitalhub/pkg/logger"
	"github.com/seregaa020292/capitalhub/pkg/utils"
)

const (
	basePrefix    = "api-market:"
	cacheDuration = 3600
)

// Market UseCase
type marketUC struct {
	cfg        *config.Config
	marketRepo market.Repository
	redisRepo  market.RedisRepository
	logger     logger.Logger
}

// Market UseCase constructor
func NewMarketUseCase(
	cfg *config.Config,
	marketRepo market.Repository,
	redisRepo market.RedisRepository,
	logger logger.Logger,
) market.UseCase {
	return &marketUC{
		cfg:        cfg,
		marketRepo: marketRepo,
		redisRepo:  redisRepo,
		logger:     logger,
	}
}

// Create market
func (useCase *marketUC) Create(ctx context.Context, market *model.Market) (*model.Market, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.Create")
	defer span.Finish()

	if err := utils.ValidateStruct(ctx, market); err != nil {
		return nil, httpErrors.NewBadRequestError(errors.WithMessage(err, "marketUC.Create.ValidateStruct"))
	}

	n, err := useCase.marketRepo.Create(ctx, market)
	if err != nil {
		return nil, err
	}

	return n, err
}

// Update market item
func (useCase *marketUC) Update(ctx context.Context, market *model.Market) (*model.Market, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.Update")
	defer span.Finish()

	updatedUser, err := useCase.marketRepo.Update(ctx, market)
	if err != nil {
		return nil, err
	}

	if err = useCase.redisRepo.DeleteMarketCtx(ctx, useCase.getKeyWithPrefix(market.MarketID.String())); err != nil {
		useCase.logger.Errorf("marketUC.Update.DeleteMarketCtx: %v", err)
	}

	return updatedUser, nil
}

// Get market by id
func (useCase *marketUC) GetByID(ctx context.Context, marketID uuid.UUID) (*model.MarketBase, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.GetByID")
	defer span.Finish()

	marketBase, err := useCase.redisRepo.GetMarketByIDCtx(ctx, useCase.getKeyWithPrefix(marketID.String()))
	if err != nil {
		useCase.logger.Errorf("marketUC.GetByID.GetByIDCtx: %v", err)
	}
	if marketBase != nil {
		return marketBase, nil
	}

	n, err := useCase.marketRepo.GetByID(ctx, marketID)
	if err != nil {
		return nil, err
	}

	if err = useCase.redisRepo.SetMarketCtx(ctx, useCase.getKeyWithPrefix(marketID.String()), cacheDuration, n); err != nil {
		useCase.logger.Errorf("marketUC.GetByID.SetByIDCtx: %s", err)
	}

	return n, nil
}

// Get market by user id
func (useCase *marketUC) GetByUserID(ctx context.Context, userID uuid.UUID) (*[]model.MarketRegister, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.GetByUserID")
	defer span.Finish()

	n, err := useCase.marketRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return n, nil
}

// Delete market
func (useCase *marketUC) Delete(ctx context.Context, marketID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.Delete")
	defer span.Finish()

	if err := useCase.marketRepo.Delete(ctx, marketID); err != nil {
		return err
	}

	if err := useCase.redisRepo.DeleteMarketCtx(ctx, useCase.getKeyWithPrefix(marketID.String())); err != nil {
		useCase.logger.Errorf("marketUC.Delete.DeleteMarketCtx: %v", err)
	}

	return nil
}

// Get markets
func (useCase *marketUC) GetAll(ctx context.Context, pq *utils.PaginationQuery) (*model.MarketList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.GetAll")
	defer span.Finish()

	return useCase.marketRepo.GetAll(ctx, pq)
}

// Find nes by title
func (useCase *marketUC) SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*model.MarketList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "marketUC.SearchByTitle")
	defer span.Finish()

	return useCase.marketRepo.SearchByTitle(ctx, title, query)
}

func (useCase *marketUC) getKeyWithPrefix(marketID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, marketID)
}
