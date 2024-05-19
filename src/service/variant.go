package service

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
)

type Variant struct {
	Repo *repository.VariantRepository
}

func (s *Variant) Update(ctx context.Context, variant database.Variant) (database.Variant, error) {
	return s.Repo.Update(ctx, variant)
}

func (s *Variant) Create(ctx context.Context, variant database.Variant) (database.Variant, error) {
	return s.Repo.Insert(ctx, variant)
}

func (s *Variant) GetAll(ctx context.Context) ([]database.Variant, error) {
	return s.Repo.List(ctx)
}

func (s *Variant) Get(ctx context.Context, id int64) (database.Variant, error) {
	return s.Repo.Get(ctx, id)
}

func NewVariantService(repo *repository.VariantRepository) Variant {
	return Variant{Repo: repo}
}
