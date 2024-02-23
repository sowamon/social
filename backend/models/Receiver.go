package models

import "gorm.io/gorm"

type Receiver struct {
	gorm.Model
	ChatId     uint
	ReceiverId uint
}
