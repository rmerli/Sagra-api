package main

import (
	"context"
	"encoding/gob"
	"gtmx/src/database"
	"gtmx/src/router"
	"log"
	"time"

	"github.com/antonlindstrom/pgstore"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	// db, err := sql.Open("postgres", "postgresql://sagra:sagra@localhost/sagra_go?sslmode=disable")
	conn, err := pgx.Connect(ctx, "postgresql://sagra:sagra@localhost/sagra_go?sslmode=disable")

	store, err := pgstore.NewPGStore("postgres://sagra:sagra@localhost/sagra_go?sslmode=disable", []byte("secret-key"))
	gob.Register(database.User{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer store.Close()

	// Run a background goroutine to clean up expired sessions from the database.
	defer store.StopCleanup(store.Cleanup(time.Minute * 5))

	queries := database.New(conn)

	if err != nil {
		println(err.Error())
		return
	}

	app := echo.New()
	router := router.New(app, queries, store)
	router.SetRoutes()

	app.Logger.Fatal(app.Start(":8080"))
}
