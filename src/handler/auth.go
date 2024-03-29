package handler

import (
	"gtmx/src/service"
	"gtmx/src/view/auth"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func (h AuthHandler) HandleSignIn(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	_, err := h.AuthService.RegisterUser(c.Request().Context(), email, password)

	if err != nil {
		return err
	}

	return nil
}

func (h AuthHandler) HandleRegister(c echo.Context) error {
	return render(c, auth.RegisterView())
}

func (h AuthHandler) HandleShowLogin(c echo.Context) error {

	session, err := session.Get("session-key", c)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, ok := session.Values["user"]

	if ok {
		return c.Redirect(http.StatusMovedPermanently, "/products")
	}

	return render(c, auth.LoginView())
}

func (h AuthHandler) HandleLogin(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer

	session, err := session.Get("session-key", c)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, ok := session.Values["user"]

	if ok {
		return c.Redirect(http.StatusMovedPermanently, "/products")
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := h.AuthService.Repository.Db.GetUser(c.Request().Context(), email)

	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}

	if user.Password != password {
		log.Println("wrong password")
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}

	session.Values["user"] = user

	if err = session.Save(r, w); err != nil {
		log.Fatalf("Error saving session: %v", err)
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}

	return c.Redirect(http.StatusMovedPermanently, "/admin/product")
}
