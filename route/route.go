package route

import (
	"echo-notes/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoute(server *echo.Echo) {
	server.GET("/api/v1/notes", controller.GetAll)
	server.GET("/api/v1/notes/:id", controller.GetByID)
	server.POST("/api/v1/notes", controller.Create)
	server.PUT("/api/v1/notes/:id", controller.Update)
	server.DELETE("/api/v1/notes/:id", controller.Delete)
}
