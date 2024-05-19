package service

import (
	"context"
	"gtmx/src/database/repository"
	"gtmx/src/model"
)

type Category struct {
	Repo *repository.CategoryRepository
}

func (s *Category) Update(ctx context.Context, category model.Category) (model.Category, error) {
	return s.Repo.Update(ctx, category)
}

func (s *Category) Insert(ctx context.Context, category model.Category) (model.Category, error) {
	return s.Repo.Insert(ctx, category)
}

func (s *Category) GetAll(ctx context.Context) ([]model.Category, error) {
	return s.Repo.List(ctx)
}

func (s *Category) Get(ctx context.Context, id int64) (model.Category, error) {
	return s.Repo.Get(ctx, id)
}

func NewCategoryService(repo *repository.CategoryRepository) Category {
	return Category{Repo: repo}
}
