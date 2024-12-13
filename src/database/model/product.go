package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	Model
	Name       string
	Abbr       string
	Price      decimal.Decimal
	CategoryID uuid.UUID
	Category   Category
}
