package routes

import (
	"gtmx/src/database"
	"gtmx/src/database/repository"
	"gtmx/src/handler"
	"gtmx/src/service"
	"net/http"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/context"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	app *echo.Echo
}

func (r *Router) SetRoutes(db *database.Queries, store *pgstore.PGStore) {
	r.app.Use(middleware.Logger())
	r.app.Use(session.Middleware(store))

	authRoutes := r.app.Group("/admin")

	authRoutes.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer context.Clear(c.Request())
			session, err := session.Get("session-key", c)
			if err != nil {
				c.Error(err)
			}

			_, ok := session.Values["user"]
			if !ok {
				return c.Redirect(http.StatusMovedPermanently, "/login")
			}

			return next(c)
		}
	})
	// r.app.Use(middleware.Recover())

	//set custom error handler for 404 and maybe 500
	errorHandler := handler.ErrorHandler{}
	r.app.HTTPErrorHandler = errorHandler.HandleError

	//set route for static files
	r.app.Static("/static", "static")
	catalogRepo := repository.CatalogRepository{Db: db}
	userRepo := repository.UserRepository{Db: db}
	authService := service.AuthService{Repository: userRepo}

	//set product routes
	productHandler := handler.ProductHandler{Repo: &catalogRepo}
	authRoutes.GET("/products", productHandler.HandleIndex)
	authRoutes.GET("/product/:id", productHandler.HandleShow)
	authRoutes.GET("/product/new", productHandler.HandleNew)
	authRoutes.POST("/product/create", productHandler.HandleCreate)

	sectionHandler := handler.SectionHandler{Repo: &catalogRepo}
	authRoutes.GET("/sections", sectionHandler.HandleIndex)
	authRoutes.GET("/section/:id", sectionHandler.HandleShow)
	authRoutes.GET("/section/new", sectionHandler.HandleNew)
	authRoutes.POST("/section/create", sectionHandler.HandleCreate)

	categoryHandler := handler.CategoryHandler{Repo: &catalogRepo}
	authRoutes.GET("/categories", categoryHandler.HandleIndex)
	authRoutes.GET("/category/:id", categoryHandler.HandleShow)
	authRoutes.GET("/category/new", categoryHandler.HandleNew)
	authRoutes.POST("/category/create", categoryHandler.HandleCreate)

	variantHandler := handler.VariantHandler{Repo: &catalogRepo}
	authRoutes.GET("/variants", variantHandler.HandleIndex)
	authRoutes.GET("/variant/:id", variantHandler.HandleShow)
	authRoutes.GET("/variant/new", variantHandler.HandleNew)
	authRoutes.POST("/variant/create", variantHandler.HandleCreate)

	authHandler := handler.AuthHandler{AuthService: authService}
	r.app.GET("/signin", authHandler.HandleRegister)
	r.app.POST("/signin", authHandler.HandleSignIn)
	r.app.GET("/login", authHandler.HandleShowLogin)
	r.app.POST("/login", authHandler.HandleLogin)
}

func New(app *echo.Echo) Router {
	return Router{app: app}
}
