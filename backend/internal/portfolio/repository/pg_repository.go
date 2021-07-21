package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/portfolio"
	"github.com/seregaa020292/capitalhub/internal/portfolio/model"
)

// Portfolio Repository
type portfolioRepo struct {
	db *sqlx.DB
}

// Portfolio Repository constructor
func NewPortfolioRepository(db *sqlx.DB) portfolio.Repository {
	return &portfolioRepo{db: db}
}

// Создание портфеля
func (repo *portfolioRepo) Create(ctx context.Context, portfolio *model.Portfolio) (*model.Portfolio, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.Create")
	defer span.Finish()

	portfolioModel := &model.Portfolio{}
	if err := repo.db.QueryRowxContext(ctx, createQuery,
		portfolio.Title, portfolio.Active, portfolio.UserID, portfolio.CurrencyID,
	).StructScan(portfolioModel); err != nil {
		return nil, errors.Wrap(err, "portfolioRepo.Create.StructScan")
	}

	return portfolioModel, nil
}

// Вернуть активный портфель
func (repo *portfolioRepo) GetActive(ctx context.Context, userID uuid.UUID) (*model.Portfolio, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.GetActive")
	defer span.Finish()

	portfolioModel := &model.Portfolio{}
	if err := repo.db.QueryRowxContext(ctx, getActiveQuery, userID).StructScan(portfolioModel); err != nil {
		return nil, errors.Wrap(err, "portfolioRepo.GetActive.QueryRowxContext")
	}
	return portfolioModel, nil
}

func (repo *portfolioRepo) CheckUserPortfolio(ctx context.Context, userID uuid.UUID, portfolioID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.CheckUserPortfolio")
	defer span.Finish()

	var exist bool
	if err := repo.db.QueryRowContext(ctx, hasPortfolio, portfolioID, userID).Scan(&exist); err != nil {
		return errors.Wrap(err, "portfolioRepo.CheckUserPortfolio.QueryRowxContext")
	}

	return nil
}
