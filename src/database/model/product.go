package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string
	Abbr       string
	Price      decimal.Decimal
	CategoryID uuid.UUID
	Category   Category
}
