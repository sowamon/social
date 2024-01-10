package db

import (
  "fmt"
  "os"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func Connect() *gorm.DB {
  fmt.Println(os.Getenv("CONNECTION_STRING"))
  db, _ := gorm.Open(mysql.Open(os.Getenv("CONNECTION_STRING")), &gorm.Config{})
  return db
}

