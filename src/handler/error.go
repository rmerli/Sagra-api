package handler

import (
	"github.com/labstack/echo/v4"
)

type ErrorHandler struct{}

func (h ErrorHandler) HandleError(err error, c echo.Context) {
	// code := http.StatusInternalServerError
	// if he, ok := err.(*echo.HTTPError); ok {
	// 	code = he.Code
	// }
	//
	// if code == http.StatusNotFound {
	// 	c.String(code, "{error: resource not found}")
	// 	return
	// }
	c.Echo().DefaultHTTPErrorHandler(err, c)
}
