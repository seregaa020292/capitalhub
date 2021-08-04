package model

import (
	"time"

	"github.com/google/uuid"
)

// Market base model
type Market struct {
	MarketID     uuid.UUID `json:"marketId" db:"market_id" validate:"omitempty,uuid"`
	Title        string    `json:"title" db:"title" validate:"required"`
	Ticker       string    `json:"ticker" db:"ticker" validate:"required"`
	Content      string    `json:"content" db:"content"`
	ImageURL     *string   `json:"imageUrl,omitempty" db:"image_url" validate:"omitempty,lte=512,url"`
	CurrencyID   uuid.UUID `json:"currencyId,omitempty" db:"currency_id" validate:"required"`
	InstrumentID uuid.UUID `json:"instrumentId,omitempty" db:"instrument_id" validate:"required"`
	CreatedAt    time.Time `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

// All market response
type MarketList struct {
	TotalCount int           `json:"totalCount"`
	TotalPages int           `json:"totalPages"`
	Page       int           `json:"page"`
	Size       int           `json:"size"`
	HasMore    bool          `json:"hasMore"`
	Markets    []*MarketBase `json:"markets"`
}

// Markets base
type MarketBase struct {
	MarketID        uuid.UUID `json:"marketId" db:"market_id" validate:"omitempty,uuid"`
	Title           string    `json:"title" db:"title" validate:"required"`
	TitleInstrument string    `json:"titleInstrument" db:"title_instrument" validate:"required"`
	DescInstrument  string    `json:"descInstrument" db:"desc_instrument" validate:"required"`
	Ticker          string    `json:"ticker" db:"ticker" validate:"required"`
	Content         *string   `json:"content" db:"content"`
	ImageURL        *string   `json:"imageUrl" db:"image_url" validate:"omitempty,lte=512,url"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty" db:"updated_at"`
}

type MarketRegister struct {
	Identify string `json:"identify" db:"identify"`
	Ticker   string `json:"ticker" db:"ticker"`
}
