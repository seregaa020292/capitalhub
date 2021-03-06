//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package provider

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/provider/model"
)

// Provider use case
type UseCase interface {
	GetByTitle(ctx context.Context, title string) (*model.Provider, error)
}
