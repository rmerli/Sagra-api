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

	"github.com/labstack/echo/v4"
)

type SectionHandler struct {
	sectionService *service.Section
}

func (h SectionHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	sections, err := h.sectionService.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.IndexSection(sections)))
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

	dbSection, err := h.sectionService.Get(c.Request().Context(), id)
	if err != nil {
		return err
	}

	viewSection, err := model.Section{}.FromDatabase(dbSection)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.ShowSection(viewSection)))
}

type updateSectionPayload struct {
	Id   int64  `param:"id"`
	Name string `form:"name"`
}

func (h SectionHandler) HandleUpdate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var paylaod updateSectionPayload
	err = c.Bind(&paylaod)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditSection(model.Section{})))
	}

	s, err := h.sectionService.Get(c.Request().Context(), paylaod.Id)
	if err != nil {
		return err
	}

	s.Name = paylaod.Name

	s, err = h.sectionService.Update(c.Request().Context(), s)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_SECTION, s.ID))
}

type editSectionPayload struct {
	Id int64 `param:"id"`
}

func (h SectionHandler) HandleEdit(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload editSectionPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditSection(model.Section{})))
	}

	dbSection, err := h.sectionService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	viewSection, err := model.Section{}.FromDatabase(dbSection)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.EditSection(viewSection)))
}

func (h SectionHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.NewSection()))
}

func (h SectionHandler) HandleCreate(c echo.Context) error {
	p := database.Section{
		Name: c.FormValue("name"),
	}

	insertedSection, err := h.sectionService.Create(c.Request().Context(), p)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%d", routes.GetPath(routes.INDEX_SECTION), insertedSection.ID)

	return c.Redirect(http.StatusMovedPermanently, endpoint)
}

func NewSectionHandler(sectionService service.Section) SectionHandler {
	return SectionHandler{
		sectionService: &sectionService,
	}
}
