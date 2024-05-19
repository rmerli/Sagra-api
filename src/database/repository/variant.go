package repository

import (
	"context"
	"gtmx/src/database"
)

type VariantRepository struct {
	db *database.Queries
}

func (r VariantRepository) List(ctx context.Context) ([]database.Variant, error) {
	variants, err := r.db.ListVariants(ctx)
	if err != nil {
		return []database.Variant{}, err

	}

	return variants, nil
}

func (r VariantRepository) Get(ctx context.Context, id int64) (database.Variant, error) {
	variant, err := r.db.GetVariant(ctx, id)

	if err != nil {
		return database.Variant{}, err

	}

	return variant, nil
}

func (r VariantRepository) Insert(ctx context.Context, variant database.Variant) (database.Variant, error) {

	insertedVariant, err := r.db.CreateVariant(ctx, database.CreateVariantParams{
		Name:  variant.Name,
		Price: variant.Price,
	})
	if err != nil {
		return database.Variant{}, err
	}

	return insertedVariant, nil
}

func (r VariantRepository) Update(ctx context.Context, variant database.Variant) (database.Variant, error) {

	variant, err := r.db.UpdateVariant(ctx, database.UpdateVariantParams{
		Name:  variant.Name,
		Price: variant.Price,
	})
	if err != nil {
		return database.Variant{}, err
	}

	return variant, nil
}

func NewVariantRepository(db *database.Queries) VariantRepository {
	return VariantRepository{db: db}
}
