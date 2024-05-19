package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type Product struct {
	Repo *repository.ProductRepository
}

func (s *Product) Update(ctx context.Context, product database.Product) (database.Product, error) {
	product, err := s.Repo.Update(ctx, product)
	if err != nil {
		return database.Product{}, err
	}
	return product, nil
}

func (s *Product) Create(ctx context.Context, product database.Product) (database.Product, error) {
	return s.Repo.Insert(ctx, product)
}

func (s *Product) GetAll(ctx context.Context) ([]database.Product, error) {
	return s.Repo.List(ctx)
}

func (s *Product) Get(ctx context.Context, id int64) (database.Product, error) {
	return s.Repo.Get(ctx, id)
}

func NewProductService(repo *repository.ProductRepository) Product {
	return Product{Repo: repo}
}
