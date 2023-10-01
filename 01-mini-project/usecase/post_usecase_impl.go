package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
)

type PostUseCaseImpl struct {
	PostRepository     repository.PostRepository
	CategoryRepository repository.CategoryRepository
}

func NewPostUseCase(postRepository repository.PostRepository, categoryRepository repository.CategoryRepository) PostUseCase {
	return &PostUseCaseImpl{
		PostRepository:     postRepository,
		CategoryRepository: categoryRepository,
	}
}

func (useCase PostUseCaseImpl) Create(payload web.PostCreateRequest, user domain.User) (web.PostResponse, error) {
	slug := helper.GenerateSlug(payload.Title)
	post := domain.Post{
		UserId:     user.Id,
		CategoryId: payload.CategoryId,
		Title:      payload.Title,
		Content:    payload.Content,
		Slug:       slug,
	}
	
	post, err := useCase.PostRepository.Save(post)
	if err != nil {
		return web.PostResponse{}, err
	}
	
	category := domain.Category{
		Id: post.CategoryId,
	}
	
	category, err = useCase.CategoryRepository.FindOne(category)
	if err != nil {
		return web.PostResponse{}, err
	}
	
	return helper.ToPostResponse(post, category, user), nil
}
