package handler

import (
	"net/http"
	"sagre/src/database/model"
	"sagre/src/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type ProductHandler struct {
	productService  *service.Product
	categoryService *service.Category
}

func (h ProductHandler) HandleIndex(c echo.Context) error {
	products, err := h.productService.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, products)
}

type showProductPayload struct {
	Id uuid.UUID `param:"id"`
}

func (h ProductHandler) HandleShow(c echo.Context) error {
	var payload showProductPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	product, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, product)
}

type createProductPayload struct {
	Name       string          `form:"name"`
	Abbr       string          `form:"abbr"`
	Price      decimal.Decimal `form:"price"`
	CategoryID uuid.UUID       `form:"category_id"`
}

func (h ProductHandler) HandleCreate(c echo.Context) error {
	var payload createProductPayload

	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	product := model.Product{
		Name:       payload.Name,
		Abbr:       payload.Abbr,
		Price:      payload.Price,
		CategoryID: payload.CategoryID,
	}

	product, err = h.productService.Create(c.Request().Context(), product)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusCreated, product)
}

type updateProductPayload struct {
	Id    uuid.UUID       `param:"id"`
	Name  string          `form:"name"`
	Abbr  string          `form:"abbr"`
	Price decimal.Decimal `form:"price"`
}

func (h ProductHandler) HandleUpdate(c echo.Context) error {
	var payload updateProductPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	product, err := h.productService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	product.Name = payload.Name
	product.Abbr = payload.Abbr
	product.Price = payload.Price

	product, err = h.productService.Update(c.Request().Context(), product)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusOK, product)
}

func NewProductHandler(service *service.Product, category *service.Category) ProductHandler {
	return ProductHandler{
		productService:  service,
		categoryService: category,
	}
}
