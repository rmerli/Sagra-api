package model

import (
	"fmt"
	"gtmx/src/database"
	"math/big"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
)

type Variant struct {
	Id    int64
	Name  string
	Price string
}

func (p Variant) FromDatabase(variant database.Variant) (Variant, error) {
	price, err := variant.Price.Float64Value()

	if err != nil {
		return Variant{}, nil
	}

	return Variant{
		Id:    variant.ID,
		Name:  variant.Name,
		Price: fmt.Sprintf("%.2f", price.Float64),
	}, nil
}

func (p Variant) ToDatabase(variant Variant) database.Variant {
	price, _ := strconv.ParseFloat(variant.Price, 64)

	return database.Variant{
		ID:    variant.Id,
		Name:  variant.Name,
		Price: pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: -2, Valid: true},
	}

}
