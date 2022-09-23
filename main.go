package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // membuat instance dari echo

	// mendaftarkan route / utk method GET
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// menjalankan HTTP server
	e.Logger.Fatal(e.Start(":1323"))
}
