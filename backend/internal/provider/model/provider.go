package model

import (
	"time"

	"github.com/google/uuid"
)

// Provider model
type Provider struct {
	ProviderID  uuid.UUID `json:"providerId" db:"provider_id" validate:"omitempty,uuid"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
