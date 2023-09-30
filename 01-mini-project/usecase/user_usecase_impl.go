package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseImpl struct {
	UserRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{UserRepository: userRepository}
}

func (useCase UserUseCaseImpl) Register(payload web.UserRegisterRequest) (web.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.UserResponse{}, err
	}
	
	user := domain.User{
		Name:     payload.Name,
		Username: payload.Username,
		Email:    payload.Email,
		Password: string(hashedPassword),
	}
	
	user, err = useCase.UserRepository.Save(user)
	if err != nil {
		return web.UserResponse{}, err
	}
	
	return helper.ToUserResponse(user), nil
}
