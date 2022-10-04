package routes

import (
	"echo-notes/controller/categories"
	"echo-notes/controller/notes"
	"echo-notes/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	AuthController     users.AuthController
	NoteController     notes.NoteController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	note := e.Group("/api/v1/notes", middleware.JWTWithConfig(cl.JWTMiddleware))

	note.GET("", cl.NoteController.GetAll)
	note.GET("/:id", cl.NoteController.GetByID)
	note.POST("", cl.NoteController.Create)
	note.PUT("/:id", cl.NoteController.Update)
	note.DELETE("/:id", cl.NoteController.Delete)
	note.POST("/:id", cl.NoteController.Restore)
	note.DELETE("/force/:id", cl.NoteController.ForceDelete)

	category := e.Group("/api/v1/categories", middleware.JWTWithConfig(cl.JWTMiddleware))

	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PUT("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)

}
