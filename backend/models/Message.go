package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Sender  uint
	Content string
	ChatId  uint
	Attach  string
}
