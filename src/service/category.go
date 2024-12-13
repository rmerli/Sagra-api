package service

import (
	"context"
	"sagre/src/database/model"
	"sagre/src/database/repository"

	"github.com/google/uuid"
)

type Category struct {
	Repo *repository.Category
}

func (c *Category) Get(ctx context.Context, id uuid.UUID) (model.Category, error) {
	return c.Repo.Get(ctx, id)
}

func (c *Category) Create(ctx context.Context, Category model.Category) (model.Category, error) {
	return c.Repo.Create(ctx, Category)
}

func (c *Category) Update(ctx context.Context, Category model.Category) (model.Category, error) {
	return c.Repo.Update(ctx, Category)
}

func (c *Category) GetAll(ctx context.Context) ([]model.Category, error) {
	return c.Repo.GetAll(ctx)
}

func NewCategoryService(repo *repository.Category) Category {
	return Category{Repo: repo}
}
