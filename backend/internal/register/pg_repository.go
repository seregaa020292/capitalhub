//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package register

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/models"
)

// Register Repository
type Repository interface {
	Create(ctx context.Context, register *models.Register) (*models.Register, error)
}
