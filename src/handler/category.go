package handler

import (
	"net/http"
	"sagre/src/database/model"
	"sagre/src/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	sectionService  *service.Section
	categoryService *service.Category
}

func (h CategoryHandler) HandleIndex(c echo.Context) error {
	categories, err := h.categoryService.GetAll(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, categories)
}

type showCategoryPayload struct {
	ID uuid.UUID `param:"id"`
}

func (h CategoryHandler) HandleShow(c echo.Context) error {
	var payload showCategoryPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	category, err := h.categoryService.Get(c.Request().Context(), payload.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, category)
}

type createCategoryPayload struct {
	Name      string    `json:"name"`
	SectionId uuid.UUID `json:"section_id"`
}

func (h CategoryHandler) HandleCreate(c echo.Context) error {
	var payload createCategoryPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	section, err := h.sectionService.Get(c.Request().Context(), payload.SectionId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	category := model.Category{
		Name:    payload.Name,
		Section: section,
	}

	insertedCategory, err := h.categoryService.Create(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusCreated, insertedCategory)
}

type updateCategoryPayload struct {
	Id        uuid.UUID `param:"id"`
	Name      string    `form:"name"`
	SectionID uuid.UUID `form:"section_id"`
}

func (h CategoryHandler) HandleUpdate(c echo.Context) error {
	var payload updateCategoryPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	category, err := h.categoryService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	section, err := h.sectionService.Get(c.Request().Context(), payload.SectionID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	category.Name = payload.Name
	category.Section = section
	updatedCategory, err := h.categoryService.Update(c.Request().Context(), category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, updatedCategory)
}

func NewCategoryHandler(sectionService *service.Section, categoryService *service.Category) CategoryHandler {
	return CategoryHandler{
		sectionService:  sectionService,
		categoryService: categoryService,
	}
}
