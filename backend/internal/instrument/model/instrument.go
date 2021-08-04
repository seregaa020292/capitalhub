package model

import (
	"time"

	"github.com/google/uuid"
)

// Instrument model
type Instrument struct {
	InstrumentID uuid.UUID `json:"instrumentId" db:"instrument_id" validate:"omitempty,uuid"`
	Title        string    `json:"title" db:"title" validate:"required"`
	Description  string    `json:"description" db:"description" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}
