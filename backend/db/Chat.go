package db

import (
	"backend/models"
)

func CreateChat(creatorId uint, participantId uint) (models.IResponse, int) {
	cn := Conn()

	if cn.Where("Creator = ? AND Participant = ?", creatorId, participantId).RowsAffected > 0 {
		return models.Response(nil, "chat already exists"), 400
	}
	c := models.Chat{Creator: creatorId, Participant: participantId}
	cn.Create(&c)

	return models.Response(c, "success"), 200
}

func GetChats(id uint) (models.IResponse, int) {
	cn := Conn()

	var chats []models.Chat

	err := cn.Where("Creator = ?", id).Find(&chats).Error

	if err != nil {
		return models.Response(nil, err.Error()), 400
	}

	if len(chats) != 0 {
		for i := 0; i < len(chats); i++ {
			chats[0].Participants = GetUserById(chats[0].Participant)
		}
	}

	return models.Response(chats, "success"), 200
}

func GetIdByUsername(username string) uint {
	cn := Conn()

	var u models.User

	cn.Where("username = ?", username).Find(&u)

	return u.ID
}

func GetUserById(id uint) models.User {
	cn := Conn()

	var u models.User

	cn.Where("id = ?", id).Find(&u)

	return u
}
