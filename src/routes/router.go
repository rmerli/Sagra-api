package routes

import (
	"database/sql"
	"gtmx/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	app *echo.Echo
}

func (r *Router) SetRoutes(db *sql.DB) {
	r.app.Use(middleware.Logger())
	r.app.Use(middleware.Recover())

	//set custom error handler for 404 and maybe 500
	errorHandler := handler.ErrorHandler{}
	r.app.HTTPErrorHandler = errorHandler.HandleError

	//set route for static files
	r.app.Static("/static", "static")

	//set product routes
	productHandler := handler.ProductHandler{Db: db}
	r.app.GET("/products", productHandler.HandleIndex)
	r.app.GET("/product/:id", productHandler.HandleShow)
	r.app.GET("/product/new", productHandler.HandleNew)
	r.app.POST("/product/create", productHandler.HandleCreate)
}

func New(app *echo.Echo) Router {
	return Router{app: app}
}
