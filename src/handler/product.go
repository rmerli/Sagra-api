package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view"
	"gtmx/src/view/layout"
	"math/big"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService *service.Product
}

func (h ProductHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	products, err := h.productService.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.IndexProduct(products)))
}

type showProductPayload struct {
	Id int64 `param:"id"`
}

func (h ProductHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload showProductPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.ShowProduct(model.Product{})))
	}

	p, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	viewProduct, err := model.Product{}.FromDatabase(p)
	return render(c, layout.ProtectedViews(user, view.ShowProduct(viewProduct)))
}

func (h ProductHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.NewProduct()))
}

type createProductPayload struct {
	Name  string  `form:"name"`
	Abbr  string  `form:"abbr"`
	Price float64 `form:"price"`
}

func (h ProductHandler) HandleCreate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload createProductPayload

	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.NewProduct()))
	}

	p := database.Product{
		Name:  c.FormValue("name"),
		Abbr:  c.FormValue("abbr"),
		Price: pgtype.Numeric{Int: big.NewInt(int64(payload.Price * 100)), Exp: -2, Valid: true},
	}

	insertedProduct, err := h.productService.Create(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath(routes.INDEX_PRODUCT), insertedProduct.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}

type editProductPayload struct {
	Id int64 `param:"id"`
}

func (h ProductHandler) HandleEdit(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload editProductPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditProduct(model.Product{})))
	}

	p, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	viewProduct, err := model.Product{}.FromDatabase(p)
	return render(c, layout.ProtectedViews(user, view.EditProduct(viewProduct)))
}

type updateProductPayload struct {
	Id    int64   `param:"id"`
	Name  string  `form:"name"`
	Abbr  string  `form:"abbr"`
	Price float64 `form:"price"`
}

func (h ProductHandler) HandleUpdate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var paylaod updateProductPayload
	err = c.Bind(&paylaod)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditProduct(model.Product{})))
	}

	p, err := h.productService.Get(c.Request().Context(), paylaod.Id)
	if err != nil {
		return err
	}

	p.Name = paylaod.Name
	p.Abbr = paylaod.Abbr
	p.Price = pgtype.Numeric{Int: big.NewInt(int64(paylaod.Price * 100)), Exp: -2, Valid: true}

	p, err = h.productService.Update(c.Request().Context(), p)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_PROUCT, p.ID))
}

func NewProductHandler(service service.Product) ProductHandler {
	return ProductHandler{productService: &service}
}
