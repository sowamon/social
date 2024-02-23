package db

import (
	"backend/models"
)

func CreateChat(creatorId uint, participantId uint) (models.IResponse, int) {
	cn := Conn()

	c := models.Chat{Creator: creatorId, Participant: participantId}
	cn.Create(&c)

	return models.Response(c, "success"), 200
}

func GetChats(id uint) (models.IResponse, int) {
	cn := Conn()

	var chats []models.Chat

	err := cn.Find(&chats).Where("id = ?", id).Error

	if err != nil {
		return models.Response(nil, err.Error()), 400
	}

	return models.Response(chats, "success"), 200
}

func GetIdByUsername(username string) uint {
	cn := Conn()

	var u models.User

	cn.Debug().Where("username = ?", username).Find(&u)

	return u.ID
}
