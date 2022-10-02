package route

import (
	"echo-notes/controller"
	"echo-notes/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {
	// routes for auth
	server.POST("/api/v1/users/register", controller.Register)
	server.POST("/api/v1/users/login", controller.Login)

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))

	privateRoutes.Use(middlewares.CheckTokenMiddleware)

	// routes for notes
	privateRoutes.GET("/api/v1/notes", controller.GetAll)
	privateRoutes.GET("/api/v1/notes/:id", controller.GetByID)
	privateRoutes.POST("/api/v1/notes", controller.Create)
	privateRoutes.PUT("/api/v1/notes/:id", controller.Update)
	privateRoutes.DELETE("/api/v1/notes/:id", controller.Delete)
	privateRoutes.POST("/api/v1/notes/:id", controller.Restore)
	privateRoutes.DELETE("/api/v1/notes/force/:id", controller.ForceDelete)

	// routes for categories
	privateRoutes.GET("/api/v1/categories", controller.GetAllCategories)
	privateRoutes.POST("/api/v1/categories", controller.CreateCategory)
	privateRoutes.PUT("/api/v1/categories/:id", controller.UpdateCategory)
	privateRoutes.DELETE("/api/v1/categories/:id", controller.DeleteCategory)

	// logout
	privateRoutes.POST("/api/v1/users/logout", controller.Logout)
}
