package repository

import (
	"context"
	"gtmx/src/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	db *gorm.DB
}

func (c *Category) Get(ctx context.Context, id uuid.UUID) (model.Category, error) {
	category := model.Category{}
	category.ID = id

	result := c.db.WithContext(ctx).First(&category)
	return category, result.Error
}

func (c *Category) Create(ctx context.Context, category model.Category) (model.Category, error) {
	result := c.db.WithContext(ctx).Create(&category)
	return category, result.Error
}

func (c *Category) Update(ctx context.Context, category model.Category) (model.Category, error) {
	result := c.db.WithContext(ctx).Save(&category)
	return category, result.Error
}

func (c *Category) GetAll(ctx context.Context) ([]model.Category, error) {
	categories := []model.Category{}
	result := c.db.WithContext(ctx).Preload("Section").Find(&categories)
	return categories, result.Error
}

func (c *Category) GetByIds(ctx context.Context, ids []uuid.UUID) ([]model.Category, error) {
	categories := []model.Category{}

	result := c.db.WithContext(ctx).Preload("Section").Find(&categories, ids)
	return categories, result.Error
}

func NewCategoryRepository(db *gorm.DB) Category {
	return Category{
		db: db,
	}
}
