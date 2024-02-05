package handler

import (
	errorview "gtmx/src/view/error_view"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorHandler struct{}

func (h ErrorHandler) HandleError(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code == http.StatusNotFound {
		// Render your custom 404 template
		render(c, errorview.NotFoundView())
		return
	}
	c.Echo().DefaultHTTPErrorHandler(err, c)
}
