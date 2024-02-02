package account

import (
	"backend/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RegisterDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func Register(c echo.Context) error {
	rq := new(RegisterDTO)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, code := db.AuthRegister(rq.Username, rq.Email, rq.Password)

	return c.JSON(code, message)
}
