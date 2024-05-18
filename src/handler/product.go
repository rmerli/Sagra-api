package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/model"
	"gtmx/src/server/routes"
	"gtmx/src/service/auth"
	"gtmx/src/view/layout"
	"gtmx/src/view/product"
	"math/big"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Repo *repository.CatalogRepository
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

func (h ProductHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}
	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}

	p, err := h.Repo.GetOneProductById(c.Request().Context(), id)
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

func (h ProductHandler) HandleCreate(c echo.Context) error {
	priceString := c.FormValue("price")
	price, err := strconv.ParseFloat(priceString, 64)

	if err != nil {
		return err

	}

	p := database.Product{
		Name:  c.FormValue("name"),
		Abbr:  c.FormValue("abbr"),
		Price: pgtype.Numeric{Int: big.NewInt(int64(price * 100)), Exp: -2, Valid: true},
	}

	insertedProduct, err := h.Repo.InsertProduct(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath("index-product"), insertedProduct.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}
