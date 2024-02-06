package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	OwnerID uint
	Content string
	Attach  string
}

type Test struct {
}
