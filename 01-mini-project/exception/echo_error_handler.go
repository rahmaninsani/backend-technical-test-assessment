package exception

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/web"
	"net/http"
)

func HTTPErrorHandler(err error, c echo.Context) {
	var httpError *echo.HTTPError
	var response *web.Response
	
	if errors.As(err, &httpError) {
		response = &web.Response{
			Code:   httpError.Code,
			Status: http.StatusText(httpError.Code),
			Data:   httpError.Message,
		}
	} else {
		response = &web.Response{
			Code:   http.StatusInternalServerError,
			Status: http.StatusText(http.StatusInternalServerError),
			Data:   err.Error(),
		}
	}
	
	if err = c.JSON(response.Code, response); err != nil {
		c.Logger().Error(err)
	}
}
