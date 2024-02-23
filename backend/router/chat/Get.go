package chat

import (
	"backend/db"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type GetChat struct{}

// @Summary Get Chats
// @Description Get Chat
// @ID getChat
// @Accept  json
// @Produce  json
// @Param username path message.SendMessage true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/chat [get]
func Get(c echo.Context) error {
	rq := new(GetChat)

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, code := db.GetChats(claims.UserId)

	return c.JSON(code, data)
}
