package server

import (
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/handler"
	customMiddleware "gtmx/src/middleware"
	"gtmx/src/server/routes"
	"gtmx/src/service/auth"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	app   *echo.Echo
	db    *database.Queries
	store *pgstore.PGStore
}

func (s *Server) Start(address string) error {
	return s.app.Start(address)
}

func (s *Server) SetRoutes() {
	s.app.Use(middleware.Logger())
	s.app.Use(session.Middleware(s.store))
	s.app.Use(middleware.Recover())

	errorHandler := handler.ErrorHandler{}
	s.app.HTTPErrorHandler = errorHandler.HandleError

	s.app.Static("/static", "static")

	catalogRepo := repository.NewCatalogRepository(s.db)
	userRepo := repository.NewUserRepository(s.db)
	authService := auth.NewAuthService(userRepo)

	authHandler := handler.AuthHandler{AuthService: authService}
	s.app.GET("/signup", authHandler.HandleShowSignUp).Name = "show-sign-up"
	s.app.POST("/signup", authHandler.HandleSignUp).Name = "sign-up"
	s.app.GET("/login", authHandler.HandleShowLogin).Name = "show-login"
	s.app.POST("/login", authHandler.HandleLogin).Name = "login"
	s.app.GET("/logout", authHandler.HandleLogout).Name = "logout"

	authenticatedRoutes := s.app.Group("/admin")
	authenticatedRoutes.Use(customMiddleware.Authenticated)
	s.app.Use(customMiddleware.ResponseHeaders)

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

	routes.SetRoutesMap(s.app.Routes())
}

func New(db *database.Queries, store *pgstore.PGStore) Server {
	return Server{
		app:   echo.New(),
		store: store,
		db:    db,
	}
}
