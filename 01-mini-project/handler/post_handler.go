package handler

import (
	"github.com/labstack/echo/v4"
)

type PostHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
