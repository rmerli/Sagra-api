package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type Category struct {
	Repo *repository.CategoryRepository
}

func (s *Category) Update(ctx context.Context, category database.Category) (database.Category, error) {
	return s.Repo.Update(ctx, category)
}

func (s *Category) Insert(ctx context.Context, category database.Category) (database.Category, error) {
	return s.Repo.Insert(ctx, category)
}

func (s *Category) GetAll(ctx context.Context) ([]database.Category, error) {
	return s.Repo.List(ctx)
}

func (s *Category) Get(ctx context.Context, id int64) (database.Category, error) {
	return s.Repo.Get(ctx, id)
}

func NewCategoryService(repo *repository.CategoryRepository) Category {
	return Category{Repo: repo}
}
