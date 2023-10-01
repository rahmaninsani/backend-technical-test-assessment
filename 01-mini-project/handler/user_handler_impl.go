package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/config"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/usecase"
	"net/http"
	"time"
)

type UserHandlerImpl struct {
	UserUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) UserHandler {
	return &UserHandlerImpl{
		UserUseCase: userUseCase,
	}
}

func (handler UserHandlerImpl) Register(c echo.Context) error {
	var payload web.UserRegisterRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	user, err := handler.UserUseCase.Register(payload)
	if err != nil {
		return err
	}
	
	userResponse := web.UserResponse{
		Name:      user.Name,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	
	response := helper.Response(http.StatusCreated, userResponse, err)
	return c.JSON(http.StatusCreated, response)
}

func (handler UserHandlerImpl) Login(c echo.Context) error {
	var payload web.UserLoginRequest
	if err := c.Bind(&payload); err != nil {
		return err
	}
	
	user, err := handler.UserUseCase.Login(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Wrong email or password")
	}
	
	c.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    user.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(config.Constant.RefreshTokenExpiresIn) * time.Minute),
		HttpOnly: true,
	})
	
	userLoginResponse := web.UserLoginResponse{
		AccessToken: user.AccessToken,
	}
	
	response := helper.Response(http.StatusOK, userLoginResponse, err)
	return c.JSON(http.StatusOK, response)
}

func (handler UserHandlerImpl) RefreshAccessToken(c echo.Context) error {
	refreshTokenCookie, err := c.Cookie("refresh_token")
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Refresh access token not found")
	}
	
	payload := web.UserRefreshAccessTokenRequest{
		RefreshToken: refreshTokenCookie.Value,
	}
	
	token, err := handler.UserUseCase.RefreshAccessToken(payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	
	response := helper.Response(http.StatusOK, token, err)
	return c.JSON(http.StatusOK, response)
}
