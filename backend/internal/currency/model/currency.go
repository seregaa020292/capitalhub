package model

import (
	"github.com/google/uuid"
	"time"
)

// Currency model
type Currency struct {
	CurrencyID  uuid.UUID `json:"currency_id" db:"currency_id" validate:"omitempty,uuid"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Description string    `json:"description" db:"description" validate:"required"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
