package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/usecase"
	"net/http"
)

type CategoryHandlerImpl struct {
	CategoryUseCase usecase.CategoryUseCase
}

func NewCategoryHandler(categoryUseCase usecase.CategoryUseCase) CategoryHandler {
	return &CategoryHandlerImpl{
		CategoryUseCase: categoryUseCase,
	}
}

func (handler CategoryHandlerImpl) Create(c echo.Context) error {
	var payload web.CategoryCreateRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	categoryResponse, err := handler.CategoryUseCase.Create(payload)
	if err != nil {
		return err
	}
	
	response := helper.Response(http.StatusCreated, categoryResponse, err)
	return c.JSON(http.StatusCreated, response)
}

func (handler CategoryHandlerImpl) FindAll(c echo.Context) error {
	categoryResponses, err := handler.CategoryUseCase.FindAll()
	if err != nil {
		return err
	}
	
	response := helper.Response(http.StatusOK, categoryResponses, err)
	return c.JSON(http.StatusOK, response)
}
