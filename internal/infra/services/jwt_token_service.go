package infra_services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenService struct{}

// implements ITokenService
var _ ITokenService = (*JwtTokenService)(nil)

func NewJwtTokenService() *JwtTokenService {
	return &JwtTokenService{}
}

func (jts *JwtTokenService) Generate(id string, secret string, expTime time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": id, "exp": time.Now().Add(expTime).Unix()})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (jts *JwtTokenService) Validate(tokenString string, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", jwt.ErrTokenInvalidId
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", jwt.ErrTokenMalformed
	}

	idValue, ok := claims["id"]
	if !ok {
		return "", jwt.ErrTokenInvalidClaims
	}

	id, ok := idValue.(string)
	if !ok {
		return "", jwt.ErrTokenInvalidClaims
	}

	return id, nil
}
