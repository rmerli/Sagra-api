package service

import (
	"context"
	"gtmx/src/database/model"
	"gtmx/src/database/repository"
)

type Menu struct {
	Repo *repository.MenuRepository
}

func (s *Menu) Update(ctx context.Context, menu model.Menu) (model.Menu, error) {
	return s.Repo.Update(ctx, menu)
}

func (s *Menu) Create(ctx context.Context, menu model.Menu) (model.Menu, error) {
	return s.Repo.Insert(ctx, menu)
}

func (s *Menu) GetAll(ctx context.Context) ([]model.Menu, error) {
	return s.Repo.List(ctx)
}

func (s *Menu) Get(ctx context.Context, id int64) (model.Menu, error) {
	return s.Repo.Get(ctx, id)
}

func NewMenuService(repo *repository.MenuRepository) Menu {
	return Menu{Repo: repo}
}
