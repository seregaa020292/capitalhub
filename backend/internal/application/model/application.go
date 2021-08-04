package model

import (
	"github.com/seregaa020292/capitalhub/internal/currency/model"
)

type Dashboard struct {
	Currencies *[]model.Currency `json:"currencies"`
}
