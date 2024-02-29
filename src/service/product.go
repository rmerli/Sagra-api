package service

import (
	"github.com/labstack/echo/v4"
)

type Product struct {
	// Db  *sql.DB
	Ctx echo.Context
}

// func (p Product) GetOneById(id int) (model.Product, error) {
// 	prod, err := model.Products(model.ProductWhere.ID.EQ(id)).One(p.Ctx.Request().Context(), p.Db)
// 	if err != nil {
// 		return model.Product{}, err
// 	}
//
// 	return *prod, nil
// }
//
// func (p Product) GetAll() (model.ProductSlice, error) {
// 	products, err := model.Products().All(p.Ctx.Request().Context(), p.Db)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	return products, nil
// }
//
// func (p Product) Insert(newProduct model.Product) (model.Product, error) {
// 	print(newProduct.Name)
// 	err := newProduct.Insert(p.Ctx.Request().Context(), p.Db, boil.Infer())
// 	if err != nil {
// 		return model.Product{}, err
// 	}
// 	newProduct.Reload(p.Ctx.Request().Context(), p.Db)
// 	return newProduct, nil
// }
//
// func (p Product) Update(updatedProduct model.Product) (model.Product, error) {
// 	// colums := map[string]interface{}{"name": updatedProduct.Name, "price": updatedProduct.Price}
//
// 	// updated, err := model.Products(model.ProductWhere.ID.EQ(updatedProduct.ID)).UpdateAll(p.Ctx.Request().Context(), p.Db, colums)
// 	updated, err := updatedProduct.Update(p.Ctx.Request().Context(), p.Db, boil.Infer())
// 	log.Println(updated)
//
// 	if err != nil {
// 		return model.Product{}, err
// 	}
// 	return updatedProduct, nil
// }
