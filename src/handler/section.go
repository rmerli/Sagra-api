package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/model"
	"gtmx/src/server/routes"
	"gtmx/src/service/auth"
	"gtmx/src/view/layout"
	"gtmx/src/view/section"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SectionHandler struct {
	Repo *repository.CatalogRepository
}

func (h SectionHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	sections, err := h.Repo.ListSections(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, section.IndexView(sections)))
}

func (h SectionHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}

	dbSection, err := h.Repo.GetOneSectionById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewSection, err := model.Section{}.FromDatabase(dbSection)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, section.ShowView(viewSection)))
}

func (h SectionHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, section.NewView()))
}

func (h SectionHandler) HandleCreate(c echo.Context) error {
	p := database.Section{
		Name: c.FormValue("name"),
	}

	insertedSection, err := h.Repo.InsertSection(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath("index-section"), insertedSection.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}
