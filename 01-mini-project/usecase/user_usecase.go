package usecase

import "github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"

type UserUseCase interface {
	Register(request web.UserRegisterRequest) (web.UserResponse, error)
}
