package db

import (
	"backend/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Conn() *gorm.DB {
	if db != nil {
		return db
	}
	db, _ = gorm.Open(mysql.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Message{}, &models.Chat{})
	return db
}
