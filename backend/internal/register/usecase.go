//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package register

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/models"
)

// Register use case
type UseCase interface {
	Create(ctx context.Context, register *models.Register) (*models.Register, error)
}
