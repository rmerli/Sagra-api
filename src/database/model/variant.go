package model

import (
	"github.com/shopspring/decimal"
)

type Variant struct {
	Model
	Name  string
	Price decimal.Decimal
}
