package main

import (
	"backend/db"
	"backend/dto"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	// v1.POST("/register", db.Register)
	v1.POST("/login", Login)

	e.Logger.Fatal(e.Start(":1881"))
}

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}

func Login(c echo.Context) error {
	cn := db.Conn()
	rq := new(dto.Login)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var u models.User
	err := cn.First(&u, "username = ?", rq.Username).Error
	if err == nil {
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(rq.Password))
		if err == nil {
			return c.JSON(http.StatusOK, models.Response(models.CreateJWT(int(u.ID)), "Successfully logged in"))
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials")
		}
	}
	return err
}
