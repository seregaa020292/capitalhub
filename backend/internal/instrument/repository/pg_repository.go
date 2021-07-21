package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/instrument"
	"github.com/seregaa020292/capitalhub/internal/models"
)

type instrumentRepo struct {
	db *sqlx.DB
}

func NewInstrumentRepository(db *sqlx.DB) instrument.Repository {
	return &instrumentRepo{db: db}
}

func (repository *instrumentRepo) GetAll(ctx context.Context) (*[]models.Instrument, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "instrumentRepo.GetAll")
	defer span.Finish()

	types := &[]models.Instrument{}
	if err := repository.db.SelectContext(ctx, types, getAll); err != nil {
		return nil, errors.Wrap(err, "instrumentRepo.GetAll.GetContext")
	}

	return types, nil
}
