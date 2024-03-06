package db

import (
	"backend/models"
	"fmt"
)

func GetUser(username string) (models.IResponse, int) {
	cn := Conn()

	var u models.User
	err := cn.First(&u, "username = ?", username).Error

	fmt.Println(err)
	if err == nil {
		return models.Response(u, "love u"), 200
	} else {
		return models.Response(nil, "Couldn't find the user"), 400
	}
}
