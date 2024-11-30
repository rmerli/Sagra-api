package handler

import (
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"gtmx/src/view"
	"gtmx/src/view/layout"

	"github.com/labstack/echo/v4"
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

func NewMenuHandler(menuService service.Menu) MenuHandler {
	return MenuHandler{
		menuService: &menuService,
	}
}
