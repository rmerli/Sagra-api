package routes

import (
	"github.com/labstack/echo/v4"
)

var customRoutes = make(map[string]string)

func SetRoutesMap(routes []*echo.Route) {
	for _, route := range routes {
		customRoutes[route.Name] = route.Path
	}
}

func GetPath(name string) string {
	path, ok := customRoutes[name]
	if !ok {
		return ""
	}

	return path
}
