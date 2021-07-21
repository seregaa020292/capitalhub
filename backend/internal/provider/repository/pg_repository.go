package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/seregaa020292/capitalhub/internal/models"
	"github.com/seregaa020292/capitalhub/internal/provider"
)

// Provider Repository
type providerRepo struct {
	db *sqlx.DB
}

// Provider Repository constructor
func NewProviderRepository(db *sqlx.DB) provider.Repository {
	return &providerRepo{db: db}
}

// GetByTitle provider
func (r *providerRepo) GetByTitle(ctx context.Context, title string) (*models.Provider, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "providerRepo.GetByTitle")
	defer span.Finish()

	providerModel := &models.Provider{}
	if err := r.db.GetContext(ctx, providerModel, getByTitle, title); err != nil {
		return nil, errors.Wrap(err, "providerRepo.GetByTitle.GetContext")
	}
	return providerModel, nil
}
