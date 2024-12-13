package service

import (
	"context"
	"gtmx/src/database/model"
	"gtmx/src/database/repository"

	"github.com/google/uuid"
)

type Menu struct {
	Repo *repository.Menu
}

func (m *Menu) Get(ctx context.Context, id uuid.UUID) (model.Menu, error) {
	return m.Repo.Get(ctx, id)
}

func (m *Menu) Create(ctx context.Context, menu model.Menu) (model.Menu, error) {
	return m.Repo.Create(ctx, menu)
}

func (m *Menu) Update(ctx context.Context, menu model.Menu, categories []model.Category) (model.Menu, error) {
	menu.Categories = make([]model.MenuCategory, len(categories))

	for i, _ := range categories {
		menu.Categories[i].CategoryID = categories[i].ID
		menu.Categories[i].MenuID = menu.ID
		menu.Categories[i].Sort = 0
	}

	return m.Repo.Update(ctx, menu)
}

func (m *Menu) GetAll(ctx context.Context) ([]model.Menu, error) {
	return m.Repo.GetAll(ctx)
}

func NewMenuService(repo *repository.Menu) Menu {
	return Menu{Repo: repo}
}
