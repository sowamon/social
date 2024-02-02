package db

import (
	"backend/dto"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Register
// @Description Register
// @ID register
// @Accept  json
// @Produce  json
// @Param username path dto.Register true "Account"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/register [post]
func Register(c echo.Context) error {
	cn := Conn()
	rq := new(dto.Register)

	if err := c.Bind(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(rq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var u models.User
	err := cn.First(&u, "username = ? or email = ?", rq.Username, rq.Email).Error

	if err == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "User already exists")
	}

	u = models.User{Username: rq.Username, Password: HashPassword(rq.Password), Email: rq.Email}
	token := models.CreateJWT(u.ID)
	cn.Create(&u)
	return c.JSON(http.StatusOK, models.Response(
		RegisterResponse(u.Username, token),
		"Account Created Successfully",
	))
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

type IRegisterResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func RegisterResponse(username string, token string) IRegisterResponse {
	return IRegisterResponse{
		Username: username,
		Token:    token,
	}
}
