package message

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type GetMessage struct {
	Reciever int `query:"reciever" validate:"required"`
	Cursor   int `query:"cursor" validate:"omitempty"`
}

func Get(c echo.Context) error {
	rq := new(GetMessage)

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)
	sender := claims.UserId

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, code := db.AuthGetMessages(sender, rq.Reciever, rq.Cursor)

	return c.JSON(code, data)
}
