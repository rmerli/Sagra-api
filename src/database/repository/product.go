package repository

import (
	"context"
	"gtmx/src/database"
)

type ProductRepository struct {
	Db *database.Queries
}

func (r ProductRepository) GetAll(ctx context.Context) ([]database.Product, error) {
	products, err := r.Db.ListProducts(ctx)
	if err != nil {
		return []database.Product{}, err

	}

	return products, nil
}

func (r ProductRepository) GetOneById(ctx context.Context, id int64) (database.Product, error) {
	product, err := r.Db.GetProduct(ctx, id)

	if err != nil {
		return database.Product{}, err

	}

	return product, nil
}

func (r ProductRepository) Insert(ctx context.Context, product database.Product) (database.Product, error) {

	insertedProduct, err := r.Db.CreateProduct(ctx, database.CreateProductParams{
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: product.Price,
	})
	if err != nil {
		return database.Product{}, err
	}

	return insertedProduct, nil
}
