package db

import (
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

// @Summary Login
// @Description Login
// @ID login
// @Accept  json
// @Produce  json
// @Param username path account.LoginDTO true "Account"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/login [post]
func AuthLogin(username string, password string) (models.IResponse, int) {
	cn := Conn()

	var u models.User
	err := cn.First(&u, "username = ?", username).Error
	if err == nil {
		err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
		if err == nil {
			return models.Response(map[string]interface{}{
				"user":  u,
				"token": models.CreateJWT(u.ID),
			}, "Successfully logged in"), 200
		} else {
			return models.Response(nil, "Wrong credentials"), 400
		}
	} else {
		return models.Response(nil, "Wrong credentials"), 400
	}
}
