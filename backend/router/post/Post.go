package post

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type PostDTO struct {
	Content string `json:"content" validate:"required"`
	Attach  string `json:"attach"`
}

func Post(c echo.Context) error {
	rq := new(PostDTO)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)

	data, code := db.Post(claims.UserId, rq.Content, rq.Attach)

	return c.JSON(code, data)
}
