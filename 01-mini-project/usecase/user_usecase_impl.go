package usecase

import (
	"github.com/google/uuid"
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
	user, err := useCase.UserRepository.FindOneByEmail(payload.Email)
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	accessTokenExpiresIn := time.Duration(config.Constant.AccessTokenExpiresIn) * time.Minute
	accessToken, err := helper.GenerateToken(&user, accessTokenExpiresIn, config.Constant.AccessTokenSecretKey)
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	refreshTokenExpiresIn := time.Duration(config.Constant.RefreshTokenExpiresIn) * time.Minute
	refreshToken, err := helper.GenerateToken(&user, refreshTokenExpiresIn, config.Constant.RefreshTokenSecretKey)
	if err != nil {
		return web.UserLoginResponse{}, err
	}
	
	return helper.ToUserLoginResponse(accessToken, refreshToken), nil
}

func (useCase UserUseCaseImpl) RefreshAccessToken(payload web.UserRefreshAccessTokenRequest) (web.UserRefreshAccessTokenResponse, error) {
	tokenClaims, err := helper.ValidateToken(payload.RefreshToken, config.Constant.RefreshTokenSecretKey)
	if err != nil {
		return web.UserRefreshAccessTokenResponse{}, err
	}
	
	userIdString := tokenClaims["user_id"].(string)
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return web.UserRefreshAccessTokenResponse{}, err
	}
	
	user, err := useCase.UserRepository.FindOneByUserId(userId)
	if err != nil {
		return web.UserRefreshAccessTokenResponse{}, err
	}
	
	accessTokenExpiresIn := time.Duration(config.Constant.AccessTokenExpiresIn) * time.Minute
	accessToken, err := helper.GenerateToken(&user, accessTokenExpiresIn, config.Constant.AccessTokenSecretKey)
	if err != nil {
		return web.UserRefreshAccessTokenResponse{}, err
	}
	
	return helper.ToUserRefreshAccessTokenResponse(accessToken), nil
}
