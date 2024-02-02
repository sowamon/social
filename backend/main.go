package main

import (
	"backend/db"
	"backend/dto"
	"backend/models"
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
	r := v1.Group("") //restricted
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	r.Use(echojwt.WithConfig(config))

	v1.GET("/ping", Pong)

	v1.POST("/register", db.Register)
	v1.POST("/login", db.Login)

	r.POST("/post", Post)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":1881"))
}

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "DURSUN ÖZBEK İSTİFA")
}

func Post(c echo.Context) error {
	cn := db.Conn()
	rq := new(dto.Post)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)

	p := models.Post{OwnerID: claims.UserId, Content: rq.Content, Attach: rq.Attach}
	cn.Create(&p)

	return c.JSON(http.StatusOK, models.Response(nil, "Successfully Posted"))
}
