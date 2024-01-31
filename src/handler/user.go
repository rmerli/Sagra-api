package handler

import (
	"gtmx/src/model"
	"gtmx/src/view/user"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {

	return render(c, user.Show(model.User{Id: 5, Email: "hello@hello.com"}))
}
