package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/handler"
)

func NewPostRouter(group *echo.Group, postHandler handler.PostHandler, middlewares []echo.MiddlewareFunc) {
	post := group.Group("/posts")
	
	post.POST("", postHandler.Create, middlewares...)
	
}
