package repository

import (
	"context"
	"database/sql"
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

func (repo *portfolioRepo) Choose(ctx context.Context, portfolioID uuid.UUID, userID uuid.UUID) (bool, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.Choose")
	defer span.Finish()

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}

	if _, err := tx.Exec(clearActiveQuery, userID); err != nil {
		tx.Rollback()
		return false, err
	}

	if _, err := tx.Exec(setActiveQuery, portfolioID, userID); err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil
}

func (repo *portfolioRepo) Edit(ctx context.Context, portfolioID uuid.UUID, change *model.PortfolioChange) (*model.Portfolio, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.Edit")
	defer span.Finish()

	portfolioModel := &model.Portfolio{}
	if err := repo.db.GetContext(ctx, portfolioModel, editQuery, change.Title, change.CurrencyID, portfolioID); err != nil {
		return nil, errors.Wrap(err, "portfolioRepo.Edit.GetContext")
	}

	return portfolioModel, nil
}

func (repo *portfolioRepo) Remove(ctx context.Context, portfolioID uuid.UUID, userID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.Remove")
	defer span.Finish()

	result, err := repo.db.ExecContext(ctx, deleteQuery, portfolioID, userID)
	if err != nil {
		return errors.WithMessage(err, "portfolioRepo Remove ExecContext")
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "portfolioRepo.Remove.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "portfolioRepo.Remove.rowsAffected")
	}

	return nil
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
	if err := repo.db.QueryRowContext(ctx, hasPortfolioUserQuery, portfolioID, userID).Scan(&exist); err != nil {
		return errors.Wrap(err, "portfolioRepo.CheckUserPortfolio.QueryRowxContext")
	}

	return nil
}

func (repo *portfolioRepo) GetAllStats(ctx context.Context, userID uuid.UUID) (*[]model.PortfolioStats, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.GetAllStats")
	defer span.Finish()

	portfolios := &[]model.PortfolioStats{}
	if err := repo.db.SelectContext(ctx, portfolios, getAllTotalQuery, userID); err != nil {
		return nil, errors.Wrap(err, "portfolioRepo.GetAllStats.SelectContext")
	}
	return portfolios, nil
}

func (repo *portfolioRepo) GetStats(ctx context.Context, portfolioID uuid.UUID) (*model.PortfolioStats, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "portfolioRepo.GetStats")
	defer span.Finish()

	portfolioStats := &model.PortfolioStats{}
	if err := repo.db.QueryRowxContext(ctx, getStatsQuery, portfolioID).StructScan(portfolioStats); err != nil {
		return nil, errors.Wrap(err, "portfolioRepo.GetStats.QueryRowxContext")
	}
	return portfolioStats, nil
}
