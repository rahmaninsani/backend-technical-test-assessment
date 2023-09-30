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