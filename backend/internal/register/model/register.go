package model

import (
	"github.com/google/uuid"
	"time"
)

// Register model
type Register struct {
	RegisterID uuid.UUID `json:"register_id" db:"register_id" validate:"omitempty,uuid"`
	Identify   string    `json:"identify" db:"identify" validate:"required"`
	ProviderID uuid.UUID `json:"provider_id,omitempty" db:"provider_id" validate:"required"`
	MarketID   uuid.UUID `json:"market_id,omitempty" db:"market_id" validate:"required"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
