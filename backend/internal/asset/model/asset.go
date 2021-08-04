package model

import (
	"github.com/google/uuid"
	"time"
)

// Asset base model
type Asset struct {
	AssetID     uuid.UUID `json:"assetId" db:"asset_id" validate:"omitempty,uuid"`
	Amount      int       `json:"amount" db:"amount" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" db:"quantity" validate:"required,gt=0"`
	PortfolioID uuid.UUID `json:"portfolioId" db:"portfolio_id" validate:"required"`
	MarketID    uuid.UUID `json:"marketId" db:"market_id" validate:"required"`
	NotationAt  time.Time `json:"notationAt" db:"notation_at"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// Base Asset response
type AssetBase struct {
	AssetID     uuid.UUID `json:"assetId" db:"asset_id" validate:"omitempty,uuid"`
	UserID      uuid.UUID `json:"userId" db:"user_id" validate:"required"`
	PortfolioID uuid.UUID `json:"portfolioId" db:"portfolio_id" validate:"required"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Ticker      string    `json:"ticker" db:"ticker" validate:"required"`
	ImageURL    *string   `json:"imageUrl" db:"image_url"`
	Amount      int       `json:"amount" db:"amount" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" db:"quantity" validate:"required,gt=0"`
	NotationAt  time.Time `json:"notationAt" db:"notation_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// All Asset response
type AssetList struct {
	TotalCount int          `json:"totalCount"`
	TotalPages int          `json:"totalPages"`
	Page       int          `json:"page"`
	Size       int          `json:"size"`
	HasMore    bool         `json:"hasMore"`
	Assets     []*AssetBase `json:"assets"`
}

// Asset add model
type AssetAdd struct {
	MarketID    uuid.UUID `json:"marketId" db:"market_id" validate:"required"`
	Amount      int       `json:"amount" db:"amount" validate:"required,gt=0"`
	Quantity    int       `json:"quantity" db:"quantity" validate:"required,gt=0"`
	PortfolioID uuid.UUID `json:"portfolioId" db:"portfolio_id" validate:"required"`
	NotationAt  time.Time `json:"notationAt" db:"notation_at"`
}

// Total Asset response
type AssetTotal struct {
	MarketID             uuid.UUID `json:"marketId" db:"market_id"`
	Title                string    `json:"title" db:"title"`
	Ticker               string    `json:"ticker" db:"ticker"`
	Identify             string    `json:"identify" db:"identify"`
	ImageURL             *string   `json:"imageUrl" db:"image_url"`
	TotalAmount          int       `json:"totalAmount" db:"total_amount"`
	TotalQuantity        int       `json:"totalQuantity" db:"total_quantity"`
	TotalCount           int       `json:"totalCount" db:"total_count"`
	AveragePurchasePrice int       `json:"averagePurchasePrice" db:"average_purchase_price"`
	FirstNotationAt      time.Time `json:"firstNotationAt" db:"first_notation_at"`
}
