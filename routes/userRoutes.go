package routes

import (
	"Praktikum/controllers"
	"os"

	jwtMid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {

	//Login
	e.POST("/login", controllers.LoginController)

	Auth := e.Group("users")
	Auth.POST("", controllers.CreateUserController)
	Auth.Use(jwtMid.JWT([]byte(os.Getenv("SECRET"))))
	Auth.GET("", controllers.GetUsersController)
	Auth.GET("/:id", controllers.GetUserController)
	Auth.PUT("/:id", controllers.UpdateUserController)
	Auth.DELETE("/:id", controllers.DeleteUserController)

}
