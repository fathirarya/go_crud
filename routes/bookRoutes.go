package routes

import (
	"Praktikum/controllers"
	"os"

	jwtMid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func InitBookRoutes(e *echo.Echo) {

	Auth := e.Group("books")
	Auth.Use(jwtMid.JWT([]byte(os.Getenv("SECRET"))))
	Auth.GET("", controllers.GetBooksController)
	Auth.GET("/:id", controllers.GetBookController)
	Auth.POST("", controllers.CreateBookController)
	Auth.PUT("/:id", controllers.UpdateBookController)
	Auth.DELETE("/:id", controllers.DeleteBookController)
}
