package model

import (
	"gtmx/src/database"

	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	Id    int64
	Name  string
	Abbr  string
	Price float64
}

func NewProduct(id int64, name string, abbr string, price pgtype.Numeric) Product {
	p, _ := price.Float64Value()
	return Product{
		Id:    id,
		Name:  name,
		Abbr:  abbr,
		Price: p.Float64,
	}
}

func NewProductList(dbProducts []database.Product) []Product {

	products := make([]Product, len(dbProducts))

	for i, product := range dbProducts {
		p, _ := product.Price.Float64Value()

		products[i] = Product{
			Id:    product.ID,
			Name:  product.Name,
			Abbr:  product.Abbr,
			Price: p.Float64,
		}
	}
	return products
}
