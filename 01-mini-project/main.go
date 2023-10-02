package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/app"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/config"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/exception"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/handler"
	customMiddleware "github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/middleware"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/router"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/usecase"
	"log"
)

func init() {
	err := config.LoadConstant()
	if err != nil {
		log.Fatalln("Failed to load environment variables\n", err.Error())
	}
}

func main() {
	db := app.NewDB()

	e := echo.New()
	e.HTTPErrorHandler = exception.HTTPErrorHandler

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}` + "\n",
		Output: e.Logger.Output(),
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	userRepository := repository.NewUserRepository(db)
	postRepository := repository.NewPostRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	tagRepository := repository.NewTagRepository(db)
	postTagRepository := repository.NewPostTagRepository(db)

	userUseCase := usecase.NewUserUseCase(userRepository, postRepository, categoryRepository)
	postUseCase := usecase.NewPostUseCase(postRepository, categoryRepository, tagRepository, postTagRepository, userRepository)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)

	userHandler := handler.NewUserHandler(userUseCase)
	postHandler := handler.NewPostHandler(postUseCase)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	jwtMiddleware := customMiddleware.JWTMiddleware(userRepository)

	api := e.Group("/api")
	router.NewUserRouter(api, userHandler, []echo.MiddlewareFunc{jwtMiddleware})
	router.NewPostRouter(api, postHandler, []echo.MiddlewareFunc{jwtMiddleware})
	router.NewCategoryRouter(api, categoryHandler, []echo.MiddlewareFunc{jwtMiddleware})

	address := fmt.Sprintf(":%s", config.Constant.AppPort)
	e.Logger.Fatal(e.Start(address))
}
