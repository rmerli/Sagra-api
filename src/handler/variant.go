package handler

import (
	"gtmx/src/database/model"
	"gtmx/src/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type VariantHandler struct {
	variantService *service.Variant
}

func (h VariantHandler) HandleIndex(c echo.Context) error {
	variants, err := h.variantService.GetAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, variants)
}

type showVariantPayload struct {
	ID uuid.UUID `param:"id"`
}

func (h VariantHandler) HandleShow(c echo.Context) error {

	var payload showVariantPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	variant, err := h.variantService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, variant)
}

type createVariantPayload struct {
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}

func (h VariantHandler) HandleCreate(c echo.Context) error {
	var payload createVariantPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	p := model.Variant{
		Name:  payload.Name,
		Price: payload.Price,
	}

	insertedVariant, err := h.variantService.Create(c.Request().Context(), p)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusCreated, insertedVariant)
}

type updateVariantPayload struct {
	ID    uuid.UUID       `param:"id"`
	Name  string          `form:"name"`
	Abbr  string          `form:"abbr"`
	Price decimal.Decimal `form:"price"`
}

func (h VariantHandler) HandleUpdate(c echo.Context) error {
	var payload updateVariantPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	variant, err := h.variantService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	variant.Name = payload.Name
	variant.Price = payload.Price

	variant, err = h.variantService.Update(c.Request().Context(), variant)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusOK, variant)
}

func NewVariantHandler(variantService *service.Variant) VariantHandler {
	return VariantHandler{
		variantService: variantService,
	}
}
