//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package currency

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/models"
)

// Currency use case
type UseCase interface {
	GetAll(ctx context.Context) (*[]models.Currency, error)
}
