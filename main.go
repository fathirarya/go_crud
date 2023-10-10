package main

import (
	"Praktikum/configs"
	"Praktikum/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	configs.LoadEnv()
	configs.InitDB()
	e.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		}))

	routes.InitUserRoutes(e)
	routes.InitBookRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
