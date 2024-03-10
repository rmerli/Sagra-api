package model

import (
	"gtmx/src/database"

	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	Id        int64
	Name      string
	SectionId int64
}

func (p Category) FromDatabase(category database.Category) (Category, error) {
	return Category{
		Id:        category.ID,
		Name:      category.Name,
		SectionId: category.SectionID.Int64,
	}, nil
}

func (p Category) ToDatabase(category Category) database.Category {

	return database.Category{
		ID:        category.Id,
		Name:      category.Name,
		SectionID: pgtype.Int8{Int64: category.SectionId, Valid: true},
	}

}
