package model

import (
	"gtmx/src/database"

	"github.com/jackc/pgx/v5/pgtype"
)

type Variant struct {
	Id    int64
	Name  string
	Price float64
}

func NewVariant(id int64, name string, price pgtype.Numeric) Variant {
	p, _ := price.Float64Value()
	return Variant{
		Id:    id,
		Name:  name,
		Price: p.Float64,
	}
}

func NewVariantList(dbVariants []database.Variant) []Variant {
	variants := make([]Variant, len(dbVariants))

	for i, variant := range dbVariants {
		p, _ := variant.Price.Float64Value()

		variants[i] = Variant{
			Id:    variant.ID,
			Name:  variant.Name,
			Price: p.Float64,
		}
	}
	return variants
}
