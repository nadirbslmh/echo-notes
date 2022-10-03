package route

import (
	CategoryController "echo-notes/controller/categories"
	NoteController "echo-notes/controller/notes"
	UserController "echo-notes/controller/users"
	"echo-notes/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {
	// routes for auth
	server.POST("/api/v1/users/register", UserController.Register)
	server.POST("/api/v1/users/login", UserController.Login)

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secretkey"),
	}))

	privateRoutes.Use(middlewares.CheckTokenMiddleware)

	// routes for notes
	privateRoutes.GET("/api/v1/notes", NoteController.GetAll)
	privateRoutes.GET("/api/v1/notes/:id", NoteController.GetByID)
	privateRoutes.POST("/api/v1/notes", NoteController.Create)
	privateRoutes.PUT("/api/v1/notes/:id", NoteController.Update)
	privateRoutes.DELETE("/api/v1/notes/:id", NoteController.Delete)
	privateRoutes.POST("/api/v1/notes/:id", NoteController.Restore)
	privateRoutes.DELETE("/api/v1/notes/force/:id", NoteController.ForceDelete)

	// routes for categories
	privateRoutes.GET("/api/v1/categories", CategoryController.GetAllCategories)
	privateRoutes.POST("/api/v1/categories", CategoryController.CreateCategory)
	privateRoutes.PUT("/api/v1/categories/:id", CategoryController.UpdateCategory)
	privateRoutes.DELETE("/api/v1/categories/:id", CategoryController.DeleteCategory)

	// logout
	privateRoutes.POST("/api/v1/users/logout", UserController.Logout)
}
