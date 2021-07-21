//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package currency

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/models"
)

// Currency Repository
type Repository interface {
	GetAll(ctx context.Context) (*[]models.Currency, error)
}
