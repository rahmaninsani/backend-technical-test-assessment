package usecase

import (
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/config"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (useCase UserUseCaseImpl) Login(payload web.UserLoginRequest) (web.UserLoginResponse, error) {
	user, err := useCase.UserRepository.FindOne(payload.Email)
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	accessTokenExpirationTime := time.Now().Add(1 * time.Hour)
	accessToken, err := helper.GenerateToken(&user, accessTokenExpirationTime, config.Constant.AccessTokenSecretKey)
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	refreshTokenExpirationTime := time.Now().Add(24 * time.Hour)
	refreshToken, err := helper.GenerateToken(&user, refreshTokenExpirationTime, config.Constant.RefreshTokenSecretKey)
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	return helper.ToUserLoginResponse(accessToken, refreshToken), nil
}
