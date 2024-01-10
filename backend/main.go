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
	v1.POST("/register", Register)

	e.Logger.Fatal(e.Start(":1881"))
}

func Pong(c echo.Context) error {
	return c.String(http.StatusOK, "Pong")
}

func Register(c echo.Context) error {
	cn := db.Connect()
	u := new(dto.User)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res := cn.Model(&models.User{}).Find(&u, "username = ? or email = ?", u.Username, u.Email)

	if res.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "user already exists")
	}

	cn.Create(&models.User{Username: u.Username, Password: HashPassword(u.Password), Email: u.Email})
	return c.JSON(http.StatusOK, "User Created Successfully")
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}
