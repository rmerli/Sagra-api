package repository

import (
	"context"
	"gtmx/src/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	db *gorm.DB
}

func (p *Product) Get(ctx context.Context, id uuid.UUID) (model.Product, error) {
	product := model.Product{ID: id}
	result := p.db.WithContext(ctx).First(&product)
	return product, result.Error
}

func (p *Product) Create(ctx context.Context, product model.Product) (model.Product, error) {
	result := p.db.WithContext(ctx).Create(&product)
	return product, result.Error
}

func (v *Product) Update(ctx context.Context, product model.Product) (model.Product, error) {
	result := v.db.WithContext(ctx).Save(&product)
	return product, result.Error
}

func (v *Product) GetAll(ctx context.Context) ([]model.Product, error) {
	products := []model.Product{}
	result := v.db.WithContext(ctx).Find(&products)
	return products, result.Error
}

func NewProductRepository(db *gorm.DB) Product {
	return Product{
		db: db,
	}
}
