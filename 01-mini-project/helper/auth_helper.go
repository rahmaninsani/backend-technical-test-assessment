package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"net/http"
	"time"
)

type Claims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateToken(user *domain.User, expirationTime time.Time, secret string) (string, error) {
	claims := &Claims{
		Name: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	
	return signedToken, nil
}

func setTokenCookie(name, value string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true
	
	c.SetCookie(cookie)
}
