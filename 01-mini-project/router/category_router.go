package router

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/handler"
)

func NewCategoryRouter(group *echo.Group, categoryHandler handler.CategoryHandler, middlewares []echo.MiddlewareFunc) {
	category := group.Group("/categories")
	
	category.POST("", categoryHandler.Create, middlewares...)
}
