package model

import (
	"github.com/google/uuid"
	"time"
)

// Currency model
type Currency struct {
	CurrencyID  uuid.UUID `json:"currencyId" db:"currency_id" validate:"omitempty,uuid"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Description string    `json:"description" db:"description" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}
