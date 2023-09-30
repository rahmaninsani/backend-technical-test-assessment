package middleware

import (
	"net/http"
	"strings"
	
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/config"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/helper"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/repository"
)

func JWTMiddleware(userRepository repository.UserRepository) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.Contains(authHeader, "Bearer") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is missing")
			}
			
			tokenString := strings.SplitN(authHeader, " ", 2)[1]
			token, err := helper.ValidateToken(tokenString, config.Constant.AccessTokenSecretKey)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}
			
			claims, ok := token.Claims.(*helper.Claims)
			if !ok || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token is not valid")
			}
			
			user, err := userRepository.FindOneByUserId(claims.UserId)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
			}
			
			c.Set("user", user)
			return next(c)
		}
	}
}
