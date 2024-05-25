package service

import (
	"context"
	"gtmx/src/database/model"
	"gtmx/src/database/repository"
)

type Product struct {
	Repo *repository.ProductRepository
}

func (s *Product) Update(ctx context.Context, product model.Product) (model.Product, error) {
	return s.Repo.Update(ctx, product)
}

func (s *Product) Create(ctx context.Context, product model.Product) (model.Product, error) {
	return s.Repo.Insert(ctx, product)
}

func (s *Product) GetAll(ctx context.Context) ([]model.Product, error) {
	return s.Repo.List(ctx)
}

func (s *Product) Get(ctx context.Context, id int64) (model.Product, error) {
	return s.Repo.Get(ctx, id)
}

func NewProductService(repo *repository.ProductRepository) Product {
	return Product{Repo: repo}
}
