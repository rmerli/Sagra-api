package handler

import (
	"database/sql"
	"gtmx/src/service"
	"gtmx/src/view/product"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Db *sql.DB
}

func (h ProductHandler) HandleIndex(c echo.Context) error {
	productService := service.Product{
		Db:  h.Db,
		Ctx: c,
	}

	products, err := productService.GetAll()

	if err != nil {
		println(err.Error())
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return render(c, product.IndexView(products))
}

func (h ProductHandler) HandleShow(c echo.Context) error {
	productService := service.Product{
		Db:  h.Db,
		Ctx: c,
	}

	idString := c.Param("id")
	id, err := strconv.Atoi(idString)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	p, err := productService.GetOneById(id)

	if err != nil {
		println(err.Error())
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return render(c, product.ShowView(p))
}

func (h ProductHandler) HandleNew(c echo.Context) error {
	return render(c, product.NewView())
}

func (h ProductHandler) HandleCreate(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/product")
}
