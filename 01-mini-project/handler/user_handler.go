package handler

import (
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	RefreshAccessToken(c echo.Context) error
	GetProfile(c echo.Context) error
	GetPostList(c echo.Context) error
}
