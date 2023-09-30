package usecase

import "github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"

type UserUseCase interface {
	Register(payload web.UserRegisterRequest) (web.UserResponse, error)
	Login(payload web.UserLoginRequest) (web.UserLoginResponse, error)
}
