package message

import (
	"backend/db"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type GetMessage struct {
	Receiver int `query:"receiver" validate:"required"`
	Cursor   int `query:"cursor" validate:"omitempty"`
}

// @Summary GetMessages
// @Description Get Message
// @ID getMessage
// @Accept  json
// @Produce  json
// @Param username path message.SendMessage true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/message [get]
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

	data, code := db.GetMessages(sender, rq.Receiver, rq.Cursor)

	fmt.Println(data.Data)

	return c.JSON(code, data)
}
