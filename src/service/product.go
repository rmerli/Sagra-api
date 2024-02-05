package service

import (
	"database/sql"
	"gtmx/src/model"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Db  *sql.DB
	Ctx echo.Context
}

func (p Product) GetOneById(id int) (model.Product, error) {
	prod, err := model.Products(model.ProductWhere.ID.EQ(id)).One(p.Ctx.Request().Context(), p.Db)
	if err != nil {
		return model.Product{}, err
	}

	return *prod, nil
}

func (p Product) GetAll() (model.ProductSlice, error) {
	products, err := model.Products().All(p.Ctx.Request().Context(), p.Db)
	if err != nil {
		return nil, err
	}

	return products, nil
}
