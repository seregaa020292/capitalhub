package models

import (
	"time"

	"github.com/google/uuid"
)

// Instrument model
type Instrument struct {
	InstrumentID uuid.UUID `json:"instrument_id" db:"instrument_id" validate:"omitempty,uuid"`
	Title        string    `json:"title" db:"title" validate:"required"`
	Description  string    `json:"description" db:"description" validate:"required"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
