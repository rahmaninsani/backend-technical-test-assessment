package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/usecase"
	"net/http"
)

type PostHandlerImpl struct {
	PostUseCase usecase.PostUseCase
}

func NewPostHandler(postUseCase usecase.PostUseCase) PostHandler {
	return &PostHandlerImpl{
		PostUseCase: postUseCase,
	}
}

func (handler PostHandlerImpl) Create(c echo.Context) error {
	var payload web.PostCreateRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	user := c.Get("user").(domain.User)
	postResponse, err := handler.PostUseCase.Create(payload, user)
	if err != nil {
		return err
	}
	
	response := helper.Response(http.StatusCreated, postResponse, err)
	return c.JSON(http.StatusCreated, response)
}

func (handler PostHandlerImpl) Update(c echo.Context) error {
	user := c.Get("user").(domain.User)
	slug := c.Param("slug")
	payload := web.PostUpdateRequest{
		Slug: slug,
	}
	
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	postResponse, err := handler.PostUseCase.Update(payload, user)
	if err != nil {
		return err
	}
	
	response := helper.Response(http.StatusOK, postResponse, err)
	return c.JSON(http.StatusOK, response)
}

func (handler PostHandlerImpl) Delete(c echo.Context) error {
	user := c.Get("user").(domain.User)
	slug := c.Param("slug")
	payload := web.PostDeleteRequest{
		Slug: slug,
	}
	
	err := handler.PostUseCase.Delete(payload, user)
	if err != nil {
		return err
	}
	
	response := helper.Response(http.StatusOK, nil, err)
	return c.JSON(http.StatusOK, response)
}
