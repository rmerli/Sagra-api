package main

import (
	"database/sql"
	"gtmx/src/routes"

	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
)

func main() {

	db, err := sql.Open("postgres", "postgresql://sagra:sagra@localhost/sagra_go?sslmode=disable")
	if err != nil {
		println(err.Error())
		return
	}

	app := echo.New()
	router := routes.New(app)
	router.SetRoutes(db)
	app.Logger.Fatal(app.Start(":8080"))
}
