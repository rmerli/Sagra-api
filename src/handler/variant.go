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

type VariantHandler struct {
	variantService *service.Variant
}

func (h VariantHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	variants, err := h.variantService.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.IndexVariant(variants)))
}

func (h VariantHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	idString := c.Param("id")

	id, err := uuid.FromBytes([]byte(idString))
	if err != nil {
		return err
	}

	variant, err := h.variantService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.ShowVariant(variant)))
}

func (h VariantHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.NewVariant()))
}

func (h VariantHandler) HandleCreate(c echo.Context) error {
	price, err := decimal.NewFromString(c.FormValue("price"))
	if err != nil {
		return err
	}

	p := model.Variant{
		Name:  c.FormValue("name"),
		Price: price,
	}

	insertedVariant, err := h.variantService.Create(c.Request().Context(), p)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_VARIANT, insertedVariant.ID))
}

type editVariantPayload struct {
	ID uuid.UUID `param:"id"`
}

func (h VariantHandler) HandleEdit(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload editVariantPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditVariant(model.Variant{})))
	}

	variant, err := h.variantService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.EditVariant(variant)))
}

type updateVariantPayload struct {
	ID    uuid.UUID       `param:"id"`
	Name  string          `form:"name"`
	Abbr  string          `form:"abbr"`
	Price decimal.Decimal `form:"price"`
}

func (h VariantHandler) HandleUpdate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload updateVariantPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditVariant(model.Variant{})))
	}

	variant, err := h.variantService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return err
	}

	variant.Name = payload.Name
	variant.Price = payload.Price

	variant, err = h.variantService.Update(c.Request().Context(), variant)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_VARIANT, payload.ID))
}

func NewVariantHandler(variantService *service.Variant) VariantHandler {
	return VariantHandler{
		variantService: variantService,
	}
}
