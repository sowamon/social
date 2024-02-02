package db

import (
	"backend/dto"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Post(c echo.Context) error {
	cn := Conn()
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
