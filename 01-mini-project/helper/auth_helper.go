package helper

import (
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
