package handler

import (
	"github.com/labstack/echo/v4"
)

type PostHandler interface {
	Create(c echo.Context) error
}
