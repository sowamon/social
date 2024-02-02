package db

import (
	"backend/dto"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	cn := Conn()
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
			return c.JSON(http.StatusOK, models.Response(map[string]interface{}{
				"user":  u,
				"token": models.CreateJWT(u.ID),
			}, "Successfully logged in"))
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials")
		}
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong credentials")
	}
}
