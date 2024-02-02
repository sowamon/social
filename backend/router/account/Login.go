package account

import (
	"backend/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	rq := new(LoginDTO)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	message, code := db.AuthLogin(rq.Username, rq.Password)

	return c.JSON(code, message)
}
