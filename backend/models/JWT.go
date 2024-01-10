package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	Id uint `json:"id"`
	jwt.RegisteredClaims
}

func CreateJWT(id int) string {
	claims := &jwtCustomClaims{
		1,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return t
}
