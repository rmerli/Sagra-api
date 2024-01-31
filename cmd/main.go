package main

import (
	"gtmx/src/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	userHandler := handler.UserHandler{}
	app.GET("/user", userHandler.HandleUserShow)
	app.Logger.Fatal(app.Start(":8080"))
}
