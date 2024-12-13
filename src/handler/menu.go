package handler

import (
	"net/http"
	"sagre/src/database/model"
	"sagre/src/service"
	"sagre/src/types"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type MenuHandler struct {
	menuService     *service.Menu
	categoryService *service.Category
}

func (h MenuHandler) HandleIndex(c echo.Context) error {
	menus, err := h.menuService.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, menus)
}

type createMenuPayload struct {
	Name  string     `json:"name"`
	Start types.Date `json:"start"`
	End   types.Date `json:"end"`
}

func (h MenuHandler) HandleCreate(c echo.Context) error {
	var payload createMenuPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	menu := model.Menu{
		Name:      payload.Name,
		StartDate: pgtype.Date{Time: time.Time(payload.Start), Valid: true},
		EndDate:   pgtype.Date{Time: time.Time(payload.End), Valid: true},
	}

	menu, err = h.menuService.Create(c.Request().Context(), menu)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusCreated, menu)
}

type updateMenuPayload struct {
	Id         uuid.UUID   `param:"id"`
	Name       string      `json:"name"`
	Start      types.Date  `json:"start"`
	End        types.Date  `json:"end"`
	Categories []uuid.UUID `json:"categories"`
}

func (h MenuHandler) HandleUpdate(c echo.Context) error {
	var payload updateMenuPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	menu, err := h.menuService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	menu.Name = payload.Name
	menu.StartDate = pgtype.Date{Time: time.Time(payload.Start), Valid: true}
	menu.EndDate = pgtype.Date{Time: time.Time(payload.End), Valid: true}

	categories, err := h.categoryService.Repo.GetByIds(c.Request().Context(), payload.Categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	menu, err = h.menuService.Update(c.Request().Context(), menu, categories)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	return c.JSON(http.StatusOK, menu)
}

type showMenuPayload struct {
	Id uuid.UUID `param:"id"`
}

func (h MenuHandler) HandleShow(c echo.Context) error {
	var payload showMenuPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	menu, err := h.menuService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, menu)
}

func NewMenuHandler(menuService *service.Menu, categoryService *service.Category) MenuHandler {
	return MenuHandler{
		menuService:     menuService,
		categoryService: categoryService,
	}
}
