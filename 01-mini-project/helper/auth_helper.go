package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rahmaninsani/backend-technical-test-assessment/01-mini-project/model/domain"
	"time"
)

type Claims struct {
	UserId uuid.UUID `json:"user_id"`
	*jwt.RegisteredClaims
}

func GenerateToken(user *domain.User, ttl time.Duration, secretKey string) (string, error) {
	claims := &Claims{
		UserId: user.Id,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(ttl)),
			NotBefore: jwt.NewNumericDate(time.Now().UTC()),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	
	return signedToken, nil
}

func ValidateToken(encodedToken string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok || method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("signing method invalid")
		}
		
		return []byte(secretKey), nil
	})
	
	if err != nil {
		return nil, err
	}
	
	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	
	return tokenClaims, nil
}
