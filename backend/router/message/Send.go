package message

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type SendMessage struct {
	Receiver int    `validate:"required"`
	Content  string `validate:"required"`
	Attach   string ``
}

// @Summary SendMessage
// @Description SendMessage
// @ID sendMessage
// @Accept  json
// @Produce  json
// @Param username path message.SendMessage true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/message [post]
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

	value, code := db.SendMessage(claims.UserId, rq.Receiver, rq.Content, rq.Attach)

	return c.JSON(code, value)
}
