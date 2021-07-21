package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/currency"
	"github.com/seregaa020292/capitalhub/internal/currency/model"
)

type currencyRepo struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) currency.Repository {
	return &currencyRepo{db: db}
}

func (repository *currencyRepo) GetAll(ctx context.Context) (*[]model.Currency, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "currencyRepo.GetAll")
	defer span.Finish()

	currencies := &[]model.Currency{}
	if err := repository.db.SelectContext(ctx, currencies, getAll); err != nil {
		return nil, errors.Wrap(err, "currencyRepo.GetAll.GetContext")
	}

	return currencies, nil
}
