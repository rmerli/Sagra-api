package handler

import (
	"gtmx/src/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view"
	"gtmx/src/view/layout"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
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

	id, err := strconv.ParseInt(idString, 10, 64)
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
	priceString := c.FormValue("price")
	price, err := strconv.ParseFloat(priceString, 64)

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

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_VARIANT, insertedVariant.Id))
}

type editVariantPayload struct {
	Id int64 `param:"id"`
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

	variant, err := h.variantService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.EditVariant(variant)))
}

type updateVariantPayload struct {
	Id    int64   `param:"id"`
	Name  string  `form:"name"`
	Abbr  string  `form:"abbr"`
	Price float64 `form:"price"`
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

	variant, err := h.variantService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	variant.Name = payload.Name
	variant.Price = payload.Price

	variant, err = h.variantService.Update(c.Request().Context(), variant)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_VARIANT, payload.Id))
}

func NewVariantHandler(variantService service.Variant) VariantHandler {
	return VariantHandler{
		variantService: &variantService,
	}
}
