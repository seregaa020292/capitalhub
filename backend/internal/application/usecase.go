//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package application

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/application/model"
)

// Application useCase interface
type UseCase interface {
	GetDashboard(ctx context.Context) (*model.Dashboard, error)
}
