package user

import (
	"backend/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetUserDto struct {
	Username string `query:"username" validate:"omitempty"`
}

func Get(c echo.Context) error {
	rq := new(GetUserDto)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, code := db.GetUser(rq.Username)

	return c.JSON(code, data)
}
