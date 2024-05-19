package repository

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/model"
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

type ProductRepository struct {
	db *database.Queries
}

func (r ProductRepository) List(ctx context.Context) ([]model.Product, error) {
	products, err := r.db.ListProducts(ctx)
	if err != nil {
		return []model.Product{}, err

	}

	return model.NewProductList(products), nil
}

func (r ProductRepository) Get(ctx context.Context, id int64) (model.Product, error) {
	product, err := r.db.GetProduct(ctx, id)

	if err != nil {
		return model.Product{}, err

	}

	return model.NewProduct(product.ID, product.Name, product.Abbr, product.Price), nil
}

func (r ProductRepository) Insert(ctx context.Context, product model.Product) (model.Product, error) {

	insertedProduct, err := r.db.CreateProduct(ctx, database.CreateProductParams{
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: pgtype.Numeric{Int: big.NewInt(int64(product.Price * 100)), Exp: -2, Valid: true},
	})

	if err != nil {
		return model.Product{}, err
	}

	product.Id = insertedProduct.ID
	return product, nil
}

func (r ProductRepository) Update(ctx context.Context, product model.Product) (model.Product, error) {
	_, err := r.db.UpdateProduct(ctx, database.UpdateProductParams{
		ID:    product.Id,
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: pgtype.Numeric{Int: big.NewInt(int64(product.Price * 100)), Exp: -2, Valid: true},
	})

	if err != nil {
		return model.Product{}, err
	}

	return product, err
}

func NewProductRepository(db *database.Queries) ProductRepository {
	return ProductRepository{db: db}
}
