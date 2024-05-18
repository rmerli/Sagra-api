package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type Product struct {
	Repo *repository.CatalogRepository
}

func (s *Product) Update(ctx context.Context, product database.Product) (database.Product, error) {
	product, err := s.Repo.UpdateProduct(ctx, product)
	if err != nil {
		return database.Product{}, err
	}
	return product, nil
}

func (s *Product) Get(ctx context.Context, id int64) (database.Product, error) {
	return s.Repo.GetOneProductById(ctx, id)
}

func NewProductService(repo *repository.CatalogRepository) Product {
	return Product{Repo: repo}
}
