package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
)

type CategoryUseCaseImpl struct {
	CategoryRepository repository.CategoryRepository
}

func NewCategoryUseCase(categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{CategoryRepository: categoryRepository}
}

func (useCase CategoryUseCaseImpl) Create(payload web.CategoryCreateRequest) (web.CategoryResponse, error) {
	category := domain.Category{
		Name: payload.Name,
	}
	
	category, err := useCase.CategoryRepository.Save(category)
	if err != nil {
		return web.CategoryResponse{}, err
	}
	
	return helper.ToCategoryResponse(category), nil
}
