//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package register

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/register/model"
)

// Register Repository
type Repository interface {
	Create(ctx context.Context, register *model.Register) (*model.Register, error)
}
