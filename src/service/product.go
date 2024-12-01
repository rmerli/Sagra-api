package service

import (
	"context"
	"gtmx/src/database/model"
	"gtmx/src/database/repository"

	"github.com/google/uuid"
)

type Product struct {
	Repo *repository.Product
}

func (p *Product) Get(ctx context.Context, id uuid.UUID) (model.Product, error) {
	return p.Repo.Get(ctx, id)
}

func (p *Product) Create(ctx context.Context, Product model.Product) (model.Product, error) {
	return p.Repo.Create(ctx, Product)
}

func (p *Product) Update(ctx context.Context, Product model.Product) (model.Product, error) {
	return p.Repo.Update(ctx, Product)
}

func (p *Product) GetAll(ctx context.Context) ([]model.Product, error) {
	return p.Repo.GetAll(ctx)
}

func NewProductService(repo *repository.Product) Product {
	return Product{Repo: repo}
}
