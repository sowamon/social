package db

import (
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

// @Summary Register
// @Description Register
// @ID register
// @Accept  json
// @Produce  json
// @Param username path account.RegisterDTO true "Account"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/register [post]
func AuthRegister(username string, email string, password string) (models.IResponse, int) {
	cn := Conn()

	var u models.User
	err := cn.First(&u, "username = ? or email = ?", username, email).Error

	if err == nil {
		return models.Response(nil, "User already exists"), 400
	}

	u = models.User{Username: username, Password: HashPassword(password), Email: email}
	token := models.CreateJWT(u.ID)
	cn.Create(&u)
	return models.Response(
		RegisterResponse(u.Username, token),
		"Account Created Successfully",
	), 200
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
