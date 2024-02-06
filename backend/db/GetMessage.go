package db

import "backend/models"

// @Summary GetMessage
// @Description Get Message
// @ID getMessage
// @Accept  json
// @Produce  json
// @Param username path message.SendMessage true "Message"
// @Header 200 {string} Token "qwerty"
// @Router /api/v1/message [get]
func AuthGetMessages(sender uint, receiver int, cursor int) (models.IResponse, int) {
	cn := Conn()

	var m []models.Message

	err := cn.Debug().Limit(10).Order("id DESC").Find(&m, "((sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)) AND id <= ?", sender, receiver, receiver, sender, cursor).Error
	if err == nil && len(m) != 0 {
		return models.Response(m, "success"), 200
	} else {
		return models.Response(nil, "no message"), 400
	}
}
