//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package instrument

import (
	"context"

	"github.com/seregaa020292/capitalhub/internal/instrument/model"
)

// Instrument Repository
type Repository interface {
	GetAll(ctx context.Context) (*[]model.Instrument, error)
}
