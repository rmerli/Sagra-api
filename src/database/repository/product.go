package repository

import (
	"context"
	"gtmx/src/database"
)

type ProductRepository struct {
	db *database.Queries
}

func (r ProductRepository) List(ctx context.Context) ([]database.Product, error) {
	products, err := r.db.ListProducts(ctx)
	if err != nil {
		return []database.Product{}, err

	}

	return products, nil
}

func (r ProductRepository) Get(ctx context.Context, id int64) (database.Product, error) {
	product, err := r.db.GetProduct(ctx, id)

	if err != nil {
		return database.Product{}, err

	}

	return product, nil
}

func (r ProductRepository) Insert(ctx context.Context, product database.Product) (database.Product, error) {

	insertedProduct, err := r.db.CreateProduct(ctx, database.CreateProductParams{
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: product.Price,
	})
	if err != nil {
		return database.Product{}, err
	}

	return insertedProduct, nil
}

func (r ProductRepository) Update(ctx context.Context, product database.Product) (database.Product, error) {
	return r.db.UpdateProduct(ctx, database.UpdateProductParams{
		ID:    product.ID,
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: product.Price,
	})
}

func NewProductRepository(db *database.Queries) ProductRepository {
	return ProductRepository{db: db}
}
