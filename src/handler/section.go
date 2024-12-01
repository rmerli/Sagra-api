package handler

import (
	"gtmx/src/database/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view"
	"gtmx/src/view/layout"
	"net/http"

	"github.com/google/uuid"
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

type showSectionPayload struct {
	ID uuid.UUID `param:"id"`
}

func (h SectionHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload showSectionPayload

	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditSection(model.Section{})))
	}

	section, err := h.sectionService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.ShowSection(section)))
}

type updateSectionPayload struct {
	ID   uuid.UUID `param:"id"`
	Name string    `form:"name"`
}

func (h SectionHandler) HandleUpdate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload updateSectionPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.EditSection(model.Section{})))
	}

	s, err := h.sectionService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return err
	}

	s.Name = payload.Name

	s, err = h.sectionService.Update(c.Request().Context(), s)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_SECTION, s.ID))
}

type editSectionPayload struct {
	ID uuid.UUID `param:"id"`
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

	section, err := h.sectionService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.EditSection(section)))
}

func (h SectionHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.NewSection()))
}

type createSectionPayload struct {
	Name string `form:"name"`
}

func (h SectionHandler) HandleCreate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload createSectionPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.NewSection()))
	}

	section, err := h.sectionService.Create(c.Request().Context(), model.Section{Name: payload.Name})
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_SECTION, section.ID))
}

func NewSectionHandler(sectionService *service.Section) SectionHandler {
	return SectionHandler{
		sectionService: sectionService,
	}
}
