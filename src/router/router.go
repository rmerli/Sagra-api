package router

import (
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/handler"
	customMiddleware "gtmx/src/middleware"
	"gtmx/src/router/routes"
	"gtmx/src/service/auth"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	app   *echo.Echo
	db    *database.Queries
	store *pgstore.PGStore
}

func (r *Router) SetRoutes() {
	r.app.Use(middleware.Logger())
	r.app.Use(session.Middleware(r.store))
	// r.app.Use(middleware.Recover()f

	errorHandler := handler.ErrorHandler{}
	r.app.HTTPErrorHandler = errorHandler.HandleError

	r.app.Static("/static", "static")

	catalogRepo := repository.NewCatalogRepository(r.db)
	userRepo := repository.NewUserRepository(r.db)
	authService := auth.NewAuthService(userRepo)

	authHandler := handler.AuthHandler{AuthService: authService}
	r.app.GET("/signup", authHandler.HandleShowSignUp).Name = "show-sign-up"
	r.app.POST("/signup", authHandler.HandleSignUp).Name = "sign-up"
	r.app.GET("/login", authHandler.HandleShowLogin).Name = "show-login"
	r.app.POST("/login", authHandler.HandleLogin).Name = "login"
	r.app.GET("/logout", authHandler.HandleLogout).Name = "logout"

	authenticatedRoutes := r.app.Group("/admin")
	authenticatedRoutes.Use(customMiddleware.Authenticated)
	r.app.Use(customMiddleware.ResponseHeaders)

	productHandler := handler.ProductHandler{Repo: &catalogRepo}
	authenticatedRoutes.GET("/products", productHandler.HandleIndex).Name = "index-product"
	authenticatedRoutes.POST("/products", productHandler.HandleCreate).Name = "create-product"
	authenticatedRoutes.GET("/products/:id", productHandler.HandleShow).Name = "show-prouct"
	authenticatedRoutes.GET("/products/new", productHandler.HandleNew).Name = "new-product"

	sectionHandler := handler.SectionHandler{Repo: &catalogRepo}
	authenticatedRoutes.GET("/sections", sectionHandler.HandleIndex).Name = "index-section"
	authenticatedRoutes.POST("/sections", sectionHandler.HandleCreate).Name = "create-section"
	authenticatedRoutes.GET("/sections/:id", sectionHandler.HandleShow).Name = "show-section"
	authenticatedRoutes.GET("/sections/new", sectionHandler.HandleNew).Name = "new-section"

	categoryHandler := handler.CategoryHandler{Repo: &catalogRepo}
	authenticatedRoutes.GET("/categories", categoryHandler.HandleIndex).Name = "index-category"
	authenticatedRoutes.POST("/categoies", categoryHandler.HandleCreate).Name = "create-category"
	authenticatedRoutes.GET("/categories/:id", categoryHandler.HandleShow).Name = "show-category"
	authenticatedRoutes.GET("/categories/new", categoryHandler.HandleNew).Name = "new-category"

	variantHandler := handler.VariantHandler{Repo: &catalogRepo}
	authenticatedRoutes.GET("/variants", variantHandler.HandleIndex).Name = "index-variant"
	authenticatedRoutes.POST("/variants", variantHandler.HandleCreate).Name = "create-variant"
	authenticatedRoutes.GET("/variants/:id", variantHandler.HandleShow).Name = "show-variant"
	authenticatedRoutes.GET("/variants/new", variantHandler.HandleNew).Name = "new-variant"

	routes.SetRoutesMap(r.app.Routes())
}

func New(app *echo.Echo, db *database.Queries, store *pgstore.PGStore) Router {
	return Router{
		app:   app,
		store: store,
		db:    db,
	}
}
