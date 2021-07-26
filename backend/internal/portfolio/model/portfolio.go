package model

import (
	"time"

	"github.com/google/uuid"

	"github.com/seregaa020292/capitalhub/internal/asset/model"
)

// Portfolio base model
type Portfolio struct {
	PortfolioID uuid.UUID `json:"portfolioId" db:"portfolio_id" validate:"omitempty,uuid"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Active      bool      `json:"active" db:"active" validate:"required"`
	UserID      uuid.UUID `json:"userId" db:"user_id" validate:"required"`
	CurrencyID  uuid.UUID `json:"currencyId" db:"currency_id" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type PortfolioList struct {
	Portfolio  Portfolio          `json:"portfolio"`
	AssetTotal []model.AssetTotal `json:"assetTotal"`
}

type PortfolioStats struct {
	PortfolioID   uuid.UUID `json:"portfolioId" db:"portfolio_id"`
	Title         string    `json:"title" db:"title"`
	Active        bool      `json:"active" db:"active"`
	CurrencyID    uuid.UUID `json:"currencyId" db:"currency_id"`
	CurrencyTitle string    `json:"currencyTitle" db:"currency_title"`
	CurrencyDesc  string    `json:"currencyDesc" db:"currency_desc"`
	AssetQuantity int       `json:"assetQuantity" db:"asset_quantity"`
	AssetAmount   int       `json:"assetAmount" db:"asset_amount"`
}
