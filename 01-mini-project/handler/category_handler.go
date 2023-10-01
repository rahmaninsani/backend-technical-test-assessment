package handler

import (
	"github.com/labstack/echo/v4"
)

type CategoryHandler interface {
	Create(c echo.Context) error
}
