package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/application"
	"github.com/seregaa020292/capitalhub/internal/application/model"
)

// Application Repository
type applicationRepo struct {
	db *sqlx.DB
}

// Application Repository constructor
func NewApplicationRepository(db *sqlx.DB) application.Repository {
	return &applicationRepo{db: db}
}

func (repository *applicationRepo) GetDashboard(ctx context.Context) (*model.Dashboard, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "applicationRepo.GetDashboard")
	defer span.Finish()

	dashboard := &model.Dashboard{}
	if err := repository.db.QueryRowxContext(ctx, getDashboard).StructScan(dashboard); err != nil {
		return nil, errors.Wrap(err, "applicationRepo.GetDashboard.QueryRowxContext")
	}

	return dashboard, nil
}
