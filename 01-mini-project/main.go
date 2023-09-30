package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
		Output: app.Logger.Output(),
	}))
	app.Use(middleware.CORS())
	app.Use(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())
	
	app.Logger.Fatal(app.Start(":1323"))
}