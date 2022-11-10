package routes

import (
	"echo-notes/app/middlewares"
	"echo-notes/controllers/categories"
	"echo-notes/controllers/notes"
	"echo-notes/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      middleware.JWTConfig
	AuthController     users.AuthController
	NoteController     notes.NoteController
	CategoryController categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	users := e.Group("/api/v1/users")

	users.POST("/register", cl.AuthController.Register)
	users.POST("/login", cl.AuthController.Login)

	note := e.Group("/api/v1/notes", middleware.JWTWithConfig(cl.JWTMiddleware))

	note.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := middlewares.GetUser(c)

			if userID == nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid token",
				})
			}

			return next(c)
		}
	})

	note.GET("", cl.NoteController.GetAll)
	note.GET("/:id", cl.NoteController.GetByID)
	note.POST("", cl.NoteController.Create)
	note.PUT("/:id", cl.NoteController.Update)
	note.DELETE("/:id", cl.NoteController.Delete)
	note.POST("/:id", cl.NoteController.Restore)
	note.DELETE("/force/:id", cl.NoteController.ForceDelete)

	category := e.Group("/api/v1/categories", middleware.JWTWithConfig(cl.JWTMiddleware))

	category.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userID := middlewares.GetUser(c)

			if userID == nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "invalid token",
				})
			}

			return next(c)
		}
	})

	category.GET("", cl.CategoryController.GetAllCategories)
	category.POST("", cl.CategoryController.CreateCategory)
	category.PUT("/:id", cl.CategoryController.UpdateCategory)
	category.DELETE("/:id", cl.CategoryController.DeleteCategory)

	auth := e.Group("/api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))

	auth.POST("/logout", cl.AuthController.Logout)

}
