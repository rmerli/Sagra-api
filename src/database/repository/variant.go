package repository

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/model"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

type VariantRepository struct {
	db *database.Queries
}

func (r VariantRepository) List(ctx context.Context) ([]model.Variant, error) {
	variants, err := r.db.ListVariants(ctx)
	if err != nil {
		return []model.Variant{}, err

	}

	return model.NewVariantList(variants), nil
}

func (r VariantRepository) Get(ctx context.Context, id int64) (model.Variant, error) {
	variant, err := r.db.GetVariant(ctx, id)

	if err != nil {
		return model.Variant{}, err

	}

	return model.NewVariant(variant.ID, variant.Name, variant.Price), nil
}

func (r VariantRepository) Insert(ctx context.Context, variant model.Variant) (model.Variant, error) {

	insertedVariant, err := r.db.CreateVariant(ctx, database.CreateVariantParams{
		Name:  variant.Name,
		Price: pgtype.Numeric{Int: big.NewInt(int64(variant.Price * 100)), Exp: -2, Valid: true},
	})
	if err != nil {
		return model.Variant{}, err
	}

	variant.Id = insertedVariant.ID

	return variant, nil
}

func (r VariantRepository) Update(ctx context.Context, variant model.Variant) (model.Variant, error) {

	_, err := r.db.UpdateVariant(ctx, database.UpdateVariantParams{
		ID:    variant.Id,
		Name:  variant.Name,
		Price: pgtype.Numeric{Int: big.NewInt(int64(variant.Price * 100)), Exp: -2, Valid: true},
	})
	if err != nil {
		return model.Variant{}, err
	}

	return variant, nil
}

func NewVariantRepository(db *database.Queries) VariantRepository {
	return VariantRepository{db: db}
}
