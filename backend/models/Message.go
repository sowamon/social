package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Sender   uint
	Content  string
	Receiver int
	Attach   string
}
