//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package currency

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/currency/model"
)

// Currency use case
type UseCase interface {
	GetAll(ctx context.Context) (*[]model.Currency, error)
}
