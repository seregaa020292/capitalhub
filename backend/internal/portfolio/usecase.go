//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package portfolio

import (
	"context"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/portfolio/model"
)

// Portfolio useCase interface
type UseCase interface {
	CreateFirst(ctx context.Context, userID uuid.UUID) (*model.Portfolio, error)
	GetActive(ctx context.Context, userID uuid.UUID) (*model.Portfolio, error)
	CheckUserPortfolio(ctx context.Context, userID uuid.UUID, portfolioID uuid.UUID) bool
}
