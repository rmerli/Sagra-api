package server

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/handler"
	customMiddleware "gtmx/src/middleware"
	"gtmx/src/server/routes"
	"gtmx/src/service"
	"gtmx/src/service/auth"
	"os"
	"os/signal"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	app         *echo.Echo
	db          *database.Queries
	store       *pgstore.PGStore
	ServerReady chan bool
}

func (s *Server) Start(address string) error {
	s.setRoutes()
	go s.app.Start(address)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return s.app.Shutdown(ctx)
}

func (s *Server) setRoutes() {
	s.app.Use(middleware.Logger())
	s.app.Use(session.Middleware(s.store))
	s.app.Use(middleware.Recover())

	errorHandler := handler.ErrorHandler{}
	s.app.HTTPErrorHandler = errorHandler.HandleError

	s.app.Static("/static", "static")

	catalogRepo := repository.NewCatalogRepository(s.db)
	sectionRepo := repository.NewSectionRepository(s.db)
	userRepo := repository.NewUserRepository(s.db)
	authService := auth.NewAuthService(userRepo)
	productService := service.NewProductService(&catalogRepo)
	sectionService := service.NewSectionService(&sectionRepo)

	authHandler := handler.AuthHandler{AuthService: authService}
	s.app.GET("/signup", authHandler.HandleShowSignUp).Name = routes.SHOW_SIGN_UP
	s.app.POST("/signup", authHandler.HandleSignUp).Name = routes.SIGN_UP
	s.app.GET("/login", authHandler.HandleShowLogin).Name = routes.SHOW_LOGIN
	s.app.POST("/login", authHandler.HandleLogin).Name = routes.LOGIN
	s.app.GET("/logout", authHandler.HandleLogout).Name = routes.LOGOUT

	authenticatedRoutes := s.app.Group("/admin")
	authenticatedRoutes.Use(customMiddleware.Authenticated)
	s.app.Use(customMiddleware.ResponseHeaders)

	productHandler := handler.ProductHandler{Repo: &catalogRepo, Service: &productService}
	authenticatedRoutes.GET("/products", productHandler.HandleIndex).Name = routes.INDEX_PRODUCT
	authenticatedRoutes.POST("/products", productHandler.HandleCreate).Name = routes.CREATE_PRODUCT
	authenticatedRoutes.GET("/products/:id", productHandler.HandleShow).Name = routes.SHOW_PROUCT
	authenticatedRoutes.GET("/products/new", productHandler.HandleNew).Name = routes.NEW_PRODUCT
	authenticatedRoutes.GET("/products/:id/edit", productHandler.HandleEdit).Name = routes.EDIT_PROUCT
	authenticatedRoutes.POST("/products/:id/update", productHandler.HandleUpdate).Name = routes.UPDATE_PROUCT

	sectionHandler := handler.SectionHandler{Repo: &sectionRepo, Service: &sectionService}
	authenticatedRoutes.GET("/sections", sectionHandler.HandleIndex).Name = routes.INDEX_SECTION
	authenticatedRoutes.POST("/sections", sectionHandler.HandleCreate).Name = routes.CREATE_SECTION
	authenticatedRoutes.GET("/sections/:id", sectionHandler.HandleShow).Name = routes.SHOW_SECTION
	authenticatedRoutes.GET("/sections/new", sectionHandler.HandleNew).Name = routes.NEW_SECTION

	categoryHandler := handler.CategoryHandler{Repo: &catalogRepo}
	authenticatedRoutes.GET("/categories", categoryHandler.HandleIndex).Name = routes.INDEX_CATEGORY
	authenticatedRoutes.POST("/categoies", categoryHandler.HandleCreate).Name = routes.CREATE_CATEGORY
	authenticatedRoutes.GET("/categories/:id", categoryHandler.HandleShow).Name = routes.SHOW_CATEGORY
	authenticatedRoutes.GET("/categories/new", categoryHandler.HandleNew).Name = routes.NEW_CATEGORY

	variantHandler := handler.VariantHandler{Repo: &catalogRepo}
	authenticatedRoutes.GET("/variants", variantHandler.HandleIndex).Name = routes.INDEX_VARIANT
	authenticatedRoutes.POST("/variants", variantHandler.HandleCreate).Name = routes.CREATE_VARIANT
	authenticatedRoutes.GET("/variants/:id", variantHandler.HandleShow).Name = routes.SHOW_VARIANT
	authenticatedRoutes.GET("/variants/new", variantHandler.HandleNew).Name = routes.NEW_VARIANT

	routes.SetRoutesMap(s.app.Routes())
}

func New(db *database.Queries, store *pgstore.PGStore, serveReady chan bool) Server {
	return Server{
		app:         echo.New(),
		store:       store,
		db:          db,
		ServerReady: serveReady,
	}
}
