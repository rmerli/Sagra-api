package repository

import (
	"context"
	"gtmx/src/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Variant struct {
	db *gorm.DB
}

func (v *Variant) Get(ctx context.Context, id uuid.UUID) (model.Variant, error) {
	variant := model.Variant{ID: id}
	result := v.db.WithContext(ctx).First(&variant)
	return variant, result.Error
}

func (v *Variant) Create(ctx context.Context, variant model.Variant) (model.Variant, error) {
	result := v.db.WithContext(ctx).Create(&variant)
	return variant, result.Error
}

func (v *Variant) Update(ctx context.Context, variant model.Variant) (model.Variant, error) {
	result := v.db.WithContext(ctx).Save(&variant)
	return variant, result.Error
}

func (v *Variant) GetAll(ctx context.Context) ([]model.Variant, error) {
	variants := []model.Variant{}
	result := v.db.WithContext(ctx).Find(&variants)
	return variants, result.Error
}

func NewVariantRepository(db *gorm.DB) Variant {
	return Variant{
		db: db,
	}
}
