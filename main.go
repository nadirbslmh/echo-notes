package main

import (
	"echo-notes/database"
	"echo-notes/middlewares"
	"echo-notes/route"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	server := echo.New()

	middlewares.LogMiddleware(server)

	route.SetupRoute(server)

	server.Logger.Fatal(server.Start(":1323"))
}
