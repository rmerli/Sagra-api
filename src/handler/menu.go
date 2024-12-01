package handler

import (
	"gtmx/src/database/model"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/types"
	"gtmx/src/view"
	"gtmx/src/view/layout"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type MenuHandler struct {
	menuService *service.Menu
}

func (h MenuHandler) HandleIndex(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	menus, err := h.menuService.GetAll(c.Request().Context())

	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.IndexMenu(menus)))
}

func (h MenuHandler) HandleNew(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.NewMenu()))
}

type createMenuPayload struct {
	Name  string     `form:"name"`
	Start types.Date `form:"start"`
	End   types.Date `form:"end"`
}

func (h MenuHandler) HandleCreate(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload createMenuPayload

	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		log.Error(err.Error())
		return render(c, layout.ProtectedViews(user, view.NewMenu()))
	}

	menu := model.Menu{
		Name:      payload.Name,
		StartDate: pgtype.Date{Time: time.Time(payload.Start), Valid: true},
		EndDate:   pgtype.Date{Time: time.Time(payload.End), Valid: true},
	}

	menu, err = h.menuService.Create(c.Request().Context(), menu)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, view.PathReplaceId(routes.SHOW_MENU, menu.ID))
}

type showMenuPayload struct {
	Id uuid.UUID `param:"id"`
}

func (h MenuHandler) HandleShow(c echo.Context) error {
	user, err := auth.GetUser(c)
	if err != nil {
		return err
	}

	var payload showMenuPayload
	err = c.Bind(&payload)
	if err != nil {
		c.Response().Status = http.StatusBadRequest
		return render(c, layout.ProtectedViews(user, view.ShowMenu(model.Menu{})))
	}

	menu, err := h.menuService.Get(c.Request().Context(), payload.Id)
	if err != nil {
		return err
	}

	return render(c, layout.ProtectedViews(user, view.ShowMenu(menu)))
}

func NewMenuHandler(menuService *service.Menu) MenuHandler {
	return MenuHandler{
		menuService: menuService,
	}
}
