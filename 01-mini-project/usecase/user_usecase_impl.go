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
	UserRepository     repository.UserRepository
	PostRepository     repository.PostRepository
	CategoryRepository repository.CategoryRepository
}

func NewUserUseCase(userRepository repository.UserRepository, postRepository repository.PostRepository,
	categoryRepository repository.CategoryRepository) UserUseCase {
	return &UserUseCaseImpl{
		UserRepository:     userRepository,
		PostRepository:     postRepository,
		CategoryRepository: categoryRepository,
	}
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
	user := domain.User{
		Email: payload.Email,
	}

	user, err := useCase.UserRepository.FindOne(user)
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

	user := domain.User{
		Id: userId,
	}

	user, err = useCase.UserRepository.FindOne(user)
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

func (useCase UserUseCaseImpl) GetProfile(payload web.UserProfileRequest) (web.UserResponse, error) {
	user := domain.User{
		Username: payload.Username,
	}

	user, err := useCase.UserRepository.FindOne(user)
	if err != nil {
		return web.UserResponse{}, err
	}

	return helper.ToUserResponse(user), nil
}

func (useCase UserUseCaseImpl) GetPostList(payload web.UserPostListRequest) ([]web.UserPostListResponse, error) {
	user := domain.User{
		Username: payload.Username,
	}

	user, err := useCase.UserRepository.FindOne(user)
	if err != nil {
		return []web.UserPostListResponse{}, err
	}

	posts, err := useCase.PostRepository.FindAll(domain.Post{
		UserId: user.Id,
	})

	if err != nil {
		return []web.UserPostListResponse{}, err
	}

	var userPostListResponses []web.UserPostListResponse
	for _, post := range posts {
		category, err := useCase.CategoryRepository.FindOne(domain.Category{Id: post.CategoryId})
		if err != nil {
			return []web.UserPostListResponse{}, err
		}

		userPostListResponses = append(userPostListResponses, helper.ToUserPostListResponse(post, category))
	}

	return userPostListResponses, nil
}
