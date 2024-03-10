package routes

import (
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	app *echo.Echo
}

func (r *Router) SetRoutes(db *database.Queries) {
	r.app.Use(middleware.Logger())
	// r.app.Use(middleware.Recover())

	//set custom error handler for 404 and maybe 500
	errorHandler := handler.ErrorHandler{}
	r.app.HTTPErrorHandler = errorHandler.HandleError

	//set route for static files
	r.app.Static("/static", "static")
	catalogRepo := repository.CatalogRepository{Db: db}

	//set product routes
	productHandler := handler.ProductHandler{Repo: &catalogRepo}
	r.app.GET("/products", productHandler.HandleIndex)
	r.app.GET("/product/:id", productHandler.HandleShow)
	r.app.GET("/product/new", productHandler.HandleNew)
	r.app.POST("/product/create", productHandler.HandleCreate)

	sectionHandler := handler.SectionHandler{Repo: &catalogRepo}
	r.app.GET("/sections", sectionHandler.HandleIndex)
	r.app.GET("/section/:id", sectionHandler.HandleShow)
	r.app.GET("/section/new", sectionHandler.HandleNew)
	r.app.POST("/section/create", sectionHandler.HandleCreate)

	categoryHandler := handler.CategoryHandler{Repo: &catalogRepo}
	r.app.GET("/categories", categoryHandler.HandleIndex)
	r.app.GET("/category/:id", categoryHandler.HandleShow)
	r.app.GET("/category/new", categoryHandler.HandleNew)
	r.app.POST("/category/create", categoryHandler.HandleCreate)

	variantHandler := handler.VariantHandler{Repo: &catalogRepo}
	r.app.GET("/variants", variantHandler.HandleIndex)
	r.app.GET("/variant/:id", variantHandler.HandleShow)
	r.app.GET("/variant/new", variantHandler.HandleNew)
	r.app.POST("/variant/create", variantHandler.HandleCreate)
}

func New(app *echo.Echo) Router {
	return Router{app: app}
}
