//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package application

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/application/model"
)

// Application repository interface
type Repository interface {
	GetDashboard(ctx context.Context) (*model.Dashboard, error)
}
