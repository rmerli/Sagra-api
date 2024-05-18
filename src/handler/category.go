package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/model"
	"gtmx/src/server/routes"
	"gtmx/src/service/auth"
	"gtmx/src/view/category"
	"gtmx/src/view/layout"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	Repo *repository.CatalogRepository
}

func (h CategoryHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	categories, err := h.Repo.ListCategories(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, category.IndexView(categories)))
}

func (h CategoryHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}

	p, err := h.Repo.GetOneCategoryById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewCategory, err := model.Category{}.FromDatabase(p)

	return render(c, layout.ProtectedViews(user, category.ShowView(viewCategory)))
}

func (h CategoryHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	sections, err := h.Repo.ListSections(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, layout.ProtectedViews(user, category.NewView(sections)))
}

func (h CategoryHandler) HandleCreate(c echo.Context) error {
	idString := c.FormValue("section_id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}

	p := database.Category{
		Name:      c.FormValue("name"),
		SectionID: pgtype.Int8{Int64: id, Valid: true},
	}

	insertedCategory, err := h.Repo.InsertCategory(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath(routes.INDEX_CATEGORY), insertedCategory.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}
