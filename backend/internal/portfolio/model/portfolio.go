package model

import (
	"github.com/seregaa020292/capitalhub/internal/models"
	"time"

	"github.com/google/uuid"
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
	Portfolio  Portfolio           `json:"portfolio"`
	AssetTotal []models.AssetTotal `json:"assetTotal"`
}
