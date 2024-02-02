package main

import (
	"backend/db"
	"backend/dto"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env")
	}

	e := echo.New()
	e.Validator = dto.Constructor(validator.New())
	v1 := e.Group("/api/v1")

	v1.GET("/ping", Pong)

	v1.POST("/register", db.Register)
	v1.POST("/login", db.Login)

	e.Logger.Fatal(e.Start(":1881"))
}

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}
