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

func (m *Menu) Create(ctx context.Context, Menu model.Menu) (model.Menu, error) {
	return m.Repo.Create(ctx, Menu)
}

func (m *Menu) Update(ctx context.Context, Menu model.Menu) (model.Menu, error) {
	return m.Repo.Update(ctx, Menu)
}

func (m *Menu) GetAll(ctx context.Context) ([]model.Menu, error) {
	return m.Repo.GetAll(ctx)
}

func NewMenuService(repo *repository.Menu) Menu {
	return Menu{Repo: repo}
}
