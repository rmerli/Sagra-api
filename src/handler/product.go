package handler

import (
	"gtmx/src/database/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view"
	"gtmx/src/view/layout"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type ProductHandler struct {
	productService  *service.Product
	categoryService *service.Category
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
	Id uuid.UUID `param:"id"`
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

	product, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.ShowProduct(product)))

}

func (h ProductHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	categories, err := h.categoryService.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.NewProduct(categories)))
}

type createProductPayload struct {
	Name       string          `form:"name"`
	Abbr       string          `form:"abbr"`
	Price      decimal.Decimal `form:"price"`
	CategoryID uuid.UUID       `form:"category_id"`
}

func (h ProductHandler) HandleCreate(c echo.Context) error {
	_, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload createProductPayload

	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return c.Redirect(http.StatusPermanentRedirect, routes.INDEX_PRODUCT)
	}

	product := model.Product{
		Name:       payload.Name,
		Abbr:       payload.Abbr,
		Price:      payload.Price,
		CategoryID: payload.CategoryID,
	}

	product, err = h.productService.Create(c.Request().Context(), product)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_PRODUCT, product.ID))
}

type editProductPayload struct {
	Id uuid.UUID `param:"id"`
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

	product, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.EditProduct(product)))
}

type updateProductPayload struct {
	Id    uuid.UUID       `param:"id"`
	Name  string          `form:"name"`
	Abbr  string          `form:"abbr"`
	Price decimal.Decimal `form:"price"`
}

func (h ProductHandler) HandleUpdate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload updateProductPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditProduct(model.Product{})))
	}

	product, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	product.Name = payload.Name
	product.Abbr = payload.Abbr
	product.Price = payload.Price

	product, err = h.productService.Update(c.Request().Context(), product)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_PRODUCT, payload.Id))
}

func NewProductHandler(service *service.Product, category *service.Category) ProductHandler {
	return ProductHandler{
		productService:  service,
		categoryService: category,
	}
}
