package usecase

import (
	"github.com/google/uuid"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
	"strings"
)

type PostUseCaseImpl struct {
	PostRepository     repository.PostRepository
	CategoryRepository repository.CategoryRepository
	TagRepository      repository.TagRepository
	PostTagRepository  repository.PostTagRepository
}

func NewPostUseCase(postRepository repository.PostRepository, categoryRepository repository.CategoryRepository,
	tagRepository repository.TagRepository, postTagRepository repository.PostTagRepository) PostUseCase {
	return &PostUseCaseImpl{
		PostRepository:     postRepository,
		CategoryRepository: categoryRepository,
		TagRepository:      tagRepository,
		PostTagRepository:  postTagRepository,
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
	
	if len(payload.Tags) > 0 {
		for index, tagString := range payload.Tags {
			tagString = strings.ToLower(tagString)
			
			tag := domain.Tag{
				Name: tagString,
			}
			
			existingTag, err := useCase.TagRepository.FindOne(tag)
			if err != nil {
				return web.PostResponse{}, err
			}
			
			if existingTag.Id == uuid.Nil {
				tag, err = useCase.TagRepository.Save(tag)
				if err != nil {
					return web.PostResponse{}, err
				}
			}
			
			postTag := domain.PostTag{
				PostId: post.Id,
				TagId:  tag.Id,
			}
			
			postTag, err = useCase.PostTagRepository.Save(postTag)
			if err != nil {
				return web.PostResponse{}, err
			}
			
			payload.Tags[index] = tag.Name
		}
	}
	
	category := domain.Category{
		Id: post.CategoryId,
	}
	
	category, err = useCase.CategoryRepository.FindOne(category)
	if err != nil {
		return web.PostResponse{}, err
	}
	
	return helper.ToPostResponse(post, category, payload.Tags, user), nil
}

func (useCase PostUseCaseImpl) Update(payload web.PostUpdateRequest, user domain.User) (web.PostResponse, error) {
	post, err := useCase.PostRepository.FindOne(domain.Post{Slug: payload.Slug})
	if err != nil || post.UserId != user.Id {
		return web.PostResponse{}, err
	}
	
	if post.Title != payload.Title {
		post.Title = payload.Title
		post.Slug = helper.GenerateSlug(payload.Title)
	}
	post.CategoryId = payload.CategoryId
	post.Content = payload.Content
	
	post, err = useCase.PostRepository.Update(post)
	if err != nil {
		return web.PostResponse{}, err
	}
	
	err = useCase.PostTagRepository.Delete(domain.PostTag{PostId: post.Id})
	if err != nil {
		return web.PostResponse{}, err
	}
	
	var tags []string
	if len(payload.Tags) > 0 {
		for _, tagString := range payload.Tags {
			tagString = strings.ToLower(tagString)
			
			tag := domain.Tag{
				Name: tagString,
			}
			
			existingTag, err := useCase.TagRepository.FindOne(tag)
			if err != nil {
				return web.PostResponse{}, err
			}
			
			if existingTag.Id == uuid.Nil {
				tag, err = useCase.TagRepository.Save(tag)
				if err != nil {
					return web.PostResponse{}, err
				}
			}
			
			postTag := domain.PostTag{
				PostId: post.Id,
				TagId:  tag.Id,
			}
			
			postTag, err = useCase.PostTagRepository.Save(postTag)
			if err != nil {
				return web.PostResponse{}, err
			}
			
			tags = append(tags, tag.Name)
		}
	}
	
	category := domain.Category{
		Id: post.CategoryId,
	}
	
	category, err = useCase.CategoryRepository.FindOne(category)
	if err != nil {
		return web.PostResponse{}, err
	}
	
	return helper.ToPostResponse(post, category, tags, user), nil
}
