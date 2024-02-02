package db

import (
	"backend/dto"
	"backend/models"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// @Summary Message
// @Description Message
// @ID message
// @Accept  json
// @Produce  json
// @Param username path dto.Message true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/message [post]
func Message(c echo.Context) error {
	cn := Conn()
	rq := new(dto.Message)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := c.Get("user").(*jwt.Token)
	claims := models.ReadJWT(user)
	m := models.Message{Sender: claims.UserId, Content: rq.Content, Reciever: rq.Reciever, Attach: rq.Attach}

	cn.Create(&m)
	return c.JSON(http.StatusOK, models.Response(
		nil,
		"Message sent successfully",
	))
}
