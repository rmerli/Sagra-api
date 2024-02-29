package main

import (
	"context"
	"gtmx/src/database"
	"gtmx/src/routes"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	// db, err := sql.Open("postgres", "postgresql://sagra:sagra@localhost/sagra_go?sslmode=disable")
	conn, err := pgx.Connect(ctx, "postgresql://sagra:sagra@localhost/sagra_go?sslmode=disable")
	queries := database.New(conn)

	if err != nil {
		println(err.Error())
		return
	}

	app := echo.New()
	router := routes.New(app)
	router.SetRoutes(queries)
	app.Logger.Fatal(app.Start(":8080"))
}
