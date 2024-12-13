package handler

import (
	"gtmx/src/database/model"
	"gtmx/src/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SectionHandler struct {
	sectionService *service.Section
}

func (h SectionHandler) HandleIndex(c echo.Context) error {
	sections, err := h.sectionService.GetAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, sections)
}

type showSectionPayload struct {
	ID uuid.UUID `param:"id"`
}

func (h SectionHandler) HandleShow(c echo.Context) error {
	var payload showSectionPayload

	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	section, err := h.sectionService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusFound)
	}

	return c.JSON(http.StatusOK, section)
}

type updateSectionPayload struct {
	ID   uuid.UUID `param:"id"`
	Name string    `json:"name"`
}

func (h SectionHandler) HandleUpdate(c echo.Context) error {
	var payload updateSectionPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	s, err := h.sectionService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	s.Name = payload.Name

	s, err = h.sectionService.Update(c.Request().Context(), s)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, s)
}

type createSectionPayload struct {
	Name string `json:"name"`
}

func (h SectionHandler) HandleCreate(c echo.Context) error {
	var payload createSectionPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	section, err := h.sectionService.Create(c.Request().Context(), model.Section{Name: payload.Name})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, section)
}

func NewSectionHandler(sectionService *service.Section) SectionHandler {
	return SectionHandler{
		sectionService: sectionService,
	}
}
