//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package provider

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/provider/model"
)

// Provider repository interface
type Repository interface {
	GetByTitle(ctx context.Context, title string) (*model.Provider, error)
}
