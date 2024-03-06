package main

import (
	"backend/dto"
	"backend/models"
	"backend/router/account"
	"backend/router/chat"
	"backend/router/message"
	"backend/router/post"
	"backend/router/user"
	"fmt"
	"net/http"
	"os"

	_ "backend/docs"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Social API
// @version 1.0
// @description Test
// @host localhost:1881
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env")
	}

	e := echo.New()
	e.Validator = dto.Constructor(validator.New())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	v1 := e.Group("/api/v1")

	v1.GET("/ping", Pong)
	v1.POST("/register", account.Register)
	v1.POST("/login", account.Login)

	//restricted
	r := v1.Group("")
	r.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}))

	r.POST("/message", message.Send)
	r.GET("/message", message.Get)

	r.POST("/post", post.Post)
	r.GET("/post", post.Get)

	r.POST("/chat", chat.Create)
	r.GET("/chat", chat.Get)

	r.GET("/user", user.Get)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1881"))
}

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "DURSUN ÖZBEK İSTİFA")
}
