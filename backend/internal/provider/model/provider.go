package model

import (
	"time"

	"github.com/google/uuid"
)

// Provider model
type Provider struct {
	ProviderID  uuid.UUID `json:"provider_id" db:"provider_id" validate:"omitempty,uuid"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
