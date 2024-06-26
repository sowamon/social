package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Creator      uint
	Participant  uint
	Participants User `gorm:"foreignKey:Participant"`
}
