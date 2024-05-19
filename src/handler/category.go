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
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	sectionService  *service.Section
	categoryService *service.Category
}

func (h CategoryHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	categories, err := h.categoryService.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.IndexCategory(categories)))
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

	p, err := h.categoryService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewCategory, err := model.Category{}.FromDatabase(p)

	return render(c, layout.ProtectedViews(user, view.ShowCategory(viewCategory)))
}

func (h CategoryHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	sections, err := h.sectionService.GetAll(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, layout.ProtectedViews(user, view.NewCategory(sections)))
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

	insertedCategory, err := h.categoryService.Insert(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath(routes.INDEX_CATEGORY), insertedCategory.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}

func NewCategoryHandler(sectionService *service.Section, categoryService *service.Category) CategoryHandler {
	return CategoryHandler{
		sectionService:  sectionService,
		categoryService: categoryService,
	}
}
