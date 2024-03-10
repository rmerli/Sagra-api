package handler

import (
	"fmt"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/model"
	"gtmx/src/view/section"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SectionHandler struct {
	Repo *repository.CatalogRepository
}

func (h SectionHandler) HandleIndex(c echo.Context) error {
	sections, err := h.Repo.ListSections(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, section.IndexView(sections))
}

func (h SectionHandler) HandleShow(c echo.Context) error {
	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return err
	}

	p, err := h.Repo.GetOneSectionById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewSection, err := model.Section{}.FromDatabase(p)

	return render(c, section.ShowView(viewSection))
}

func (h SectionHandler) HandleNew(c echo.Context) error {
	return render(c, section.NewView())
}

func (h SectionHandler) HandleCreate(c echo.Context) error {
	p := database.Section{
		Name: c.FormValue("name"),
	}

	insertedSection, err := h.Repo.InsertSection(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("/section/%d", insertedSection.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}
