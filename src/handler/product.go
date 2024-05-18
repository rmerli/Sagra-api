package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view/layout"
	"gtmx/src/view/product"
	"math/big"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service *service.Product
	Repo    *repository.CatalogRepository
}

func (h ProductHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	products, err := h.Repo.ListProducts(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, product.IndexView(products)))
}

type showProductRequest struct {
	Id int64 `param:"id"`
}

func (h ProductHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var request showProductRequest
	err = c.Bind(&request)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, product.ShowView(model.Product{})))
	}

	p, err := h.Repo.GetOneProductById(c.Request().Context(), request.Id)
	if err != nil {
		return err
	}

	viewProduct, err := model.Product{}.FromDatabase(p)
	return render(c, layout.ProtectedViews(user, product.ShowView(viewProduct)))
}

func (h ProductHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, product.NewView()))
}

type createProductRequest struct {
	Name  string  `form:"name"`
	Abbr  string  `form:"abbr"`
	Price float64 `form:"price"`
}

func (h ProductHandler) HandleCreate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var request createProductRequest

	err = c.Bind(&request)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, product.NewView()))
	}

	p := database.Product{
		Name:  c.FormValue("name"),
		Abbr:  c.FormValue("abbr"),
		Price: pgtype.Numeric{Int: big.NewInt(int64(request.Price * 100)), Exp: -2, Valid: true},
	}

	insertedProduct, err := h.Repo.InsertProduct(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath(routes.INDEX_PRODUCT), insertedProduct.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}

type editProductRequest struct {
	Id int64 `param:"id"`
}

func (h ProductHandler) HandleEdit(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var request editProductRequest
	err = c.Bind(&request)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, product.EditView(model.Product{})))
	}

	p, err := h.Repo.GetOneProductById(c.Request().Context(), request.Id)
	if err != nil {
		return err
	}

	viewProduct, err := model.Product{}.FromDatabase(p)
	return render(c, layout.ProtectedViews(user, product.EditView(viewProduct)))
}

type updateProductRequest struct {
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

	var request updateProductRequest
	err = c.Bind(&request)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, product.EditView(model.Product{})))
	}

	p, err := h.Service.Get(c.Request().Context(), request.Id)
	if err != nil {
		return err
	}

	p.Name = request.Name
	p.Abbr = request.Abbr
	p.Price = pgtype.Numeric{Int: big.NewInt(int64(request.Price * 100)), Exp: -2, Valid: true}

	p, err = h.Service.Update(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := strings.Replace(routes.GetPath(routes.SHOW_PROUCT), ":id", fmt.Sprint(p.ID), 1)
	return c.Redirect(http.StatusMovedPermanently, endpoint)
}
