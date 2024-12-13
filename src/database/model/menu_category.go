package model

import "github.com/google/uuid"

type MenuCategory struct {
	Model
	CategoryID uuid.UUID `json:"categoryId"`
	Category   Category  `json:"category"`
	MenuID     uuid.UUID `json:"menuId"`
	Menu       Menu      `json:"menu"`
	Sort       int       `json:"sort"`
	Products   []Product `gorm:"many2many:menu_category_products;" json:"products"`
}
