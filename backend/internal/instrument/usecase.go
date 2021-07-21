//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package instrument

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/models"
)

// Instrument use case
type UseCase interface {
	GetAll(ctx context.Context) (*[]models.Instrument, error)
}
