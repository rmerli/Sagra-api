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
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
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

	p, err := h.variantService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewVariant, err := model.Variant{}.FromDatabase(p)

	return render(c, layout.ProtectedViews(user, view.ShowVariant(viewVariant)))
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

	p := database.Variant{
		Name:  c.FormValue("name"),
		Price: pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: -2, Valid: true},
	}

	insertedVariant, err := h.variantService.Create(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath(routes.INDEX_VARIANT), insertedVariant.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}

func NewVariantHandler(variantService service.Variant) VariantHandler {
	return VariantHandler{
		variantService: &variantService,
	}
}
