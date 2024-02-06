package post

import (
	"backend/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetPostDTO struct {
	Cursor int `query:"cursor" validate:"omitempty"`
}

func Get(c echo.Context) error {
	rq := new(GetPostDTO)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, code := db.Get(rq.Cursor)

	return c.JSON(code, data)
}
