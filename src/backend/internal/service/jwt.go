package service

import (
	"fundz/internal/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(id string) (string, error) {

	jwtSecret := config.Env.Jwt.JWTSECRET

	claims := JWTClaims{
		UserId: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
