package middleware

import (
	"github.com/gorilla/context"
	"github.com/labstack/echo/v4"
)

func Authenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer context.Clear(c.Request())

		// session, err := session.Get("session-key", c)
		// if err != nil {
		// 	c.Error(err)
		// }
		//
		// _, ok := session.Values["user"]
		// if !ok {
		// 	return c.Redirect(http.StatusMovedPermanently, routes.GetPath(routes.SHOW_LOGIN))
		// }

		return next(c)
	}
}

func ResponseHeaders(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer context.Clear(c.Request())
		c.Response().Header().Set(echo.HeaderCacheControl, "no-cache")
		return next(c)
	}
}
