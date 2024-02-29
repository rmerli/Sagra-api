package model

import (
	"fmt"
	"gtmx/src/database"
	"math/big"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type Product struct {
	Id    int64
	Name  string
	Abbr  string
	Price string
}

func (p Product) FromDatabase(product database.Product) (Product, error) {
	price, err := product.Price.Float64Value()

	if err != nil {
		return Product{}, nil
	}

	return Product{
		Id:    product.ID,
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: fmt.Sprintf("%.2f", price.Float64),
	}, nil
}

func (p Product) ToDatabase(product Product) database.Product {
	price, _ := strconv.ParseFloat(product.Price, 64)

	return database.Product{
		ID:    product.Id,
		Name:  product.Name,
		Abbr:  product.Abbr,
		Price: pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: -2, Valid: true},
	}

}
