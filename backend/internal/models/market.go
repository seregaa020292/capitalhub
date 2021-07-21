package models

import (
	"time"

	"github.com/google/uuid"
)

// Market base model
type Market struct {
	MarketID     uuid.UUID `json:"market_id" db:"market_id" validate:"omitempty,uuid"`
	Title        string    `json:"title" db:"title" validate:"required"`
	Ticker       string    `json:"ticker" db:"ticker" validate:"required"`
	Content      string    `json:"content" db:"content"`
	ImageURL     *string   `json:"image_url,omitempty" db:"image_url" validate:"omitempty,lte=512,url"`
	CurrencyID   uuid.UUID `json:"currency_id,omitempty" db:"currency_id" validate:"required"`
	InstrumentID uuid.UUID `json:"instrument_id,omitempty" db:"instrument_id" validate:"required"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// All market response
type MarketList struct {
	TotalCount int           `json:"total_count"`
	TotalPages int           `json:"total_pages"`
	Page       int           `json:"page"`
	Size       int           `json:"size"`
	HasMore    bool          `json:"has_more"`
	Markets    []*MarketBase `json:"markets"`
}

// Markets base
type MarketBase struct {
	MarketID        uuid.UUID `json:"market_id" db:"market_id" validate:"omitempty,uuid"`
	Title           string    `json:"title" db:"title" validate:"required"`
	TitleInstrument string    `json:"title_instrument" db:"title_instrument" validate:"required"`
	DescInstrument  string    `json:"desc_instrument" db:"desc_instrument" validate:"required"`
	Ticker          string    `json:"ticker" db:"ticker" validate:"required"`
	Content         *string   `json:"content" db:"content"`
	ImageURL        *string   `json:"image_url" db:"image_url" validate:"omitempty,lte=512,url"`
	UpdatedAt       time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type MarketRegister struct {
	Identify string `json:"identify" db:"identify"`
	Ticker   string `json:"ticker" db:"ticker"`
}
