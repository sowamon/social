package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	OwnerID uint   `json:"owner_id"`
	Content string `json:"content"`
	Attach  string `json:"attach"`
}

type Test struct {
}
