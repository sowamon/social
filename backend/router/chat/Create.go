package chat

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type CreateChat struct {
	Participant string `validate:"required"`
}

// @Summary SendMessage
// @Description SendMessage
// @ID sendMessage
// @Accept  json
// @Produce  json
// @Param username path message.SendMessage true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/message [post]
func Create(c echo.Context) error {
	rq := new(CreateChat)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)

	value, code := db.CreateChat(claims.UserId, db.GetIdByUsername(rq.Participant))

	return c.JSON(code, value)
}
