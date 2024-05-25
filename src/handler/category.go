package handler

import (
	"gtmx/src/database/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view"
	"gtmx/src/view/layout"
	"net/http"
	"strconv"

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

	category, err := h.categoryService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.ShowCategory(category)))
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

	section, err := h.sectionService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	p := model.Category{
		Name:    c.FormValue("name"),
		Section: section,
	}

	insertedCategory, err := h.categoryService.Insert(c.Request().Context(), p)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_CATEGORY, insertedCategory.Id))
}

type editCategoryPayload struct {
	Id int64 `param:"id"`
}

func (h CategoryHandler) HandleEdit(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload editCategoryPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditCategory(model.Category{}, nil)))
	}

	category, err := h.categoryService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	sections, err := h.sectionService.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.EditCategory(category, sections)))
}

type updateCategoryPayload struct {
	Id        int64  `param:"id"`
	Name      string `form:"name"`
	SectionID int64  `form:"section_id"`
}

func (h CategoryHandler) HandleUpdate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload updateCategoryPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditCategory(model.Category{}, nil)))
	}

	category, err := h.categoryService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	section, err := h.sectionService.Get(c.Request().Context(), payload.SectionID)
	if err != nil {
		return err
	}

	category.Name = payload.Name
	category.Section = section

	category, err = h.categoryService.Update(c.Request().Context(), category)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_CATEGORY, category.Id))
}

func NewCategoryHandler(sectionService *service.Section, categoryService *service.Category) CategoryHandler {
	return CategoryHandler{
		sectionService:  sectionService,
		categoryService: categoryService,
	}
}
