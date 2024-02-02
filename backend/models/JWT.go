package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId uint `json:"id"`
	jwt.RegisteredClaims
}

func CreateJWT(id uint) string {
	claims := &JwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return t
}

func IsExpired(token jwt.Token) *jwt.NumericDate {
	claims := token.Claims.(*JwtCustomClaims)
	status := claims.ExpiresAt
	return status
}

func ReadJWT(token *jwt.Token) *JwtCustomClaims {
	return token.Claims.(*JwtCustomClaims)
}
