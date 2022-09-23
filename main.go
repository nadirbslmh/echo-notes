package main

import (
	"echo-notes/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	e := echo.New() // membuat instance dari echo

	// mendaftarkan route / utk method GET
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "alterra")
	})

	// menjalankan HTTP server
	e.Logger.Fatal(e.Start(":1323"))
}
