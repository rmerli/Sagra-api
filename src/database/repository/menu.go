package repository

import (
	"context"
	"gtmx/src/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	db *gorm.DB
}

func (m *Menu) Get(ctx context.Context, id uuid.UUID) (model.Menu, error) {
	menu := model.Menu{ID: id}
	result := m.db.WithContext(ctx).First(&menu)
	return menu, result.Error
}

func (m *Menu) Create(ctx context.Context, menu model.Menu) (model.Menu, error) {
	result := m.db.WithContext(ctx).Create(&menu)
	return menu, result.Error
}

func (m *Menu) Update(ctx context.Context, menu model.Menu) (model.Menu, error) {
	result := m.db.WithContext(ctx).Save(&menu)
	return menu, result.Error
}

func (m *Menu) GetAll(ctx context.Context) ([]model.Menu, error) {
	menus := []model.Menu{}
	result := m.db.WithContext(ctx).Find(&menus)
	return menus, result.Error
}

func NewMenuRepository(db *gorm.DB) Menu {
	return Menu{
		db: db,
	}
}
