package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/register"
	"github.com/seregaa020292/capitalhub/internal/register/model"
)

type registerRepo struct {
	db *sqlx.DB
}

func NewRegisterRepository(db *sqlx.DB) register.Repository {
	return &registerRepo{db: db}
}

func (repository *registerRepo) Create(ctx context.Context, register *model.Register) (*model.Register, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "registerRepo.Create")
	defer span.Finish()

	var model model.Register
	if err := repository.db.QueryRowxContext(
		ctx,
		create,
		&register.Identify,
		&register.ProviderID,
		&register.MarketID,
	).StructScan(&model); err != nil {
		return nil, errors.Wrap(err, "registerRepo.Create.QueryRowxContext")
	}

	return &model, nil
}
