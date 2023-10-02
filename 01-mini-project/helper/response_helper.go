package helper

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"net/http"
)

func Response(code int, data interface{}, err error) *web.Response {
	response := &web.Response{
		Code:   code,
		Status: http.StatusText(code),
	}

	if data != nil {
		response.Data = data
	}

	if err != nil {
		response.Message = err.Error()
	}

	return response
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserLoginResponse(accessToken, refreshToken string) web.UserLoginResponse {
	return web.UserLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func ToUserRefreshAccessTokenResponse(accessToken string) web.UserRefreshAccessTokenResponse {
	return web.UserRefreshAccessTokenResponse{
		AccessToken: accessToken,
	}
}

func ToPostResponse(post domain.Post, category domain.Category, tags []string, user domain.User) web.PostResponse {
	return web.PostResponse{
		Title:   post.Title,
		Content: post.Content,
		Slug:    post.Slug,
		Category: web.PostCategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		},
		Tags: tags,
		Author: web.PostAuthorResponse{
			Name:     user.Name,
			Username: user.Username,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse

	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ToUserPostListResponse(post domain.Post, category domain.Category) web.UserPostListResponse {
	return web.UserPostListResponse{
		Title: post.Title,
		Slug:  post.Slug,
		Category: web.PostCategoryResponse{
			Id:   category.Id,
			Name: category.Name,
		},
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}
