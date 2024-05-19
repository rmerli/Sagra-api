package service

import (
	"context"
	"gtmx/src/database/repository"
	"gtmx/src/model"
)

type Variant struct {
	Repo *repository.VariantRepository
}

func (s *Variant) Update(ctx context.Context, variant model.Variant) (model.Variant, error) {
	return s.Repo.Update(ctx, variant)
}

func (s *Variant) Create(ctx context.Context, variant model.Variant) (model.Variant, error) {
	return s.Repo.Insert(ctx, variant)
}

func (s *Variant) GetAll(ctx context.Context) ([]model.Variant, error) {
	return s.Repo.List(ctx)
}

func (s *Variant) Get(ctx context.Context, id int64) (model.Variant, error) {
	return s.Repo.Get(ctx, id)
}

func NewVariantService(repo *repository.VariantRepository) Variant {
	return Variant{Repo: repo}
}
