package message

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type SendMessage struct {
	Reciever int    `json:"reciever" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Attach   string `json:"attach"`
}

func Send(c echo.Context) error {
	rq := new(SendMessage)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)

	value, code := db.AuthSendMessage(claims.UserId, rq.Reciever, rq.Content, rq.Attach)

	return c.JSON(code, value)
}
