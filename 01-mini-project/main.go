package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/config"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/exception"
	"log"
)

func init() {
	err := config.LoadConstant()
	if err != nil {
		log.Fatalln("Failed to load environment variables\n", err.Error())
	}
}

func main() {
	app := echo.New()
	app.HTTPErrorHandler = exception.HTTPErrorHandler
	
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
		Output: app.Logger.Output(),
	}))
	app.Use(middleware.CORS())
	app.Use(middleware.RemoveTrailingSlash())
	app.Use(middleware.Recover())
	
	address := fmt.Sprintf(":%s", config.Constant.AppPort)
	app.Logger.Fatal(app.Start(address))
}
