package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name  string
	Price decimal.Decimal
}
