package service

import (
	"context"
	"sagre/src/database/model"
	"sagre/src/database/repository"

	"github.com/google/uuid"
)

type Variant struct {
	Repo *repository.Variant
}

func (v *Variant) Get(ctx context.Context, id uuid.UUID) (model.Variant, error) {
	return v.Repo.Get(ctx, id)
}

func (v *Variant) Create(ctx context.Context, Variant model.Variant) (model.Variant, error) {
	return v.Repo.Create(ctx, Variant)
}

func (v *Variant) Update(ctx context.Context, Variant model.Variant) (model.Variant, error) {
	return v.Repo.Update(ctx, Variant)
}

func (v *Variant) GetAll(ctx context.Context) ([]model.Variant, error) {
	return v.Repo.GetAll(ctx)
}

func NewVariantService(repo *repository.Variant) Variant {
	return Variant{Repo: repo}
}
