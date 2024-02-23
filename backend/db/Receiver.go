package db

import (
	"backend/models"
)

func CreateReceiver(chatId uint, receiverId uint) {
	cn := Conn()

	d := models.Receiver{ChatId: chatId, ReceiverId: receiverId}
	cn.Create(&d)
}

func GetReceivers(chatId uint) []models.Receiver {
	cn := Conn()

	var receivers []models.Receiver

	cn.Find(&receivers).Where("ChatId = ?", chatId)

	return receivers
}
