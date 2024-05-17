package handler

import (
	"fmt"
	"gtmx/src/router/routes"
	authentication "gtmx/src/service/auth"
	"gtmx/src/validator"
	"gtmx/src/view/auth"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthService authentication.AuthService
}

func (h AuthHandler) HandleSignUp(c echo.Context) error {
	bag := validator.SignUp{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	err := bag.Validate(c.Request().Context(), h.AuthService.GetRepository())

	if err != nil {
		return render(c, auth.RegisterView(bag))
	}

	_, err = h.AuthService.RegisterUser(c.Request().Context(), bag.Email, bag.Password)

	if err != nil {
		return err
	}

	//fix error when ridirect it automatically logs in
	return c.Redirect(http.StatusTemporaryRedirect, routes.GetPath("login"))
}

func (h AuthHandler) HandleShowSignUp(c echo.Context) error {
	return render(c, auth.RegisterView(validator.SignUp{}))
}

func (h AuthHandler) HandleShowLogin(c echo.Context) error {

	session, err := session.Get("session-key", c)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, ok := session.Values["user"]

	if ok {
		return c.Redirect(http.StatusMovedPermanently, routes.GetPath("index-product"))
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
		return c.Redirect(http.StatusMovedPermanently, routes.GetPath("index-product"))
	}

	email := c.FormValue("email")
	password := c.FormValue("password")

	user, err := h.AuthService.GetRepository().GetUser(c.Request().Context(), email)

	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusMovedPermanently, routes.GetPath("login"))
	}

	check := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(fmt.Sprintf("%s%s", password, user.Salt)))

	if check != nil {
		log.Println("wrong password")
		return c.Redirect(http.StatusMovedPermanently, routes.GetPath("login"))
	}

	session.Values["user"] = user

	if err = session.Save(r, w); err != nil {
		log.Fatalf("Error saving session: %v", err)
		return c.Redirect(http.StatusMovedPermanently, routes.GetPath("login"))
	}

	return c.Redirect(http.StatusMovedPermanently, routes.GetPath("index-product"))
}

func (h AuthHandler) HandleLogout(c echo.Context) error {
	r := c.Request()
	w := c.Response().Writer

	session, err := session.Get("session-key", c)
	if err != nil {
		return err
	}
	session.Values["user"] = nil
	session.Options.MaxAge = -1

	if err = session.Save(r, w); err != nil {
		log.Fatalf("Error saving session: %v", err)
	}

	return c.Redirect(http.StatusMovedPermanently, routes.GetPath("login"))
}
