package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Sender   uint   `json:"sender_id"`
	Content  string `json:"content"`
	Reciever int    `json:"reciever"`
	Attach   string `json:"attach"`
}
