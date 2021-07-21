//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package portfolio

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/portfolio/model"
)

// Portfolio repository interface
type Repository interface {
	Create(ctx context.Context, portfolio *model.Portfolio) (*model.Portfolio, error)
	GetActive(ctx context.Context, userID uuid.UUID) (*model.Portfolio, error)
	CheckUserPortfolio(ctx context.Context, userID uuid.UUID, portfolioID uuid.UUID) error
}
