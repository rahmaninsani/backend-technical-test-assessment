package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/handler"
)

func NewUserRouter(group *echo.Group, userHandler handler.UserHandler, middlewares []echo.MiddlewareFunc) {
	user := group.Group("/users")
	
	user.POST("", userHandler.Register)
	user.POST("/login", userHandler.Login)
	user.GET("/refresh", userHandler.RefreshAccessToken)
	
	user.GET("/:username", userHandler.GetProfile)
}
