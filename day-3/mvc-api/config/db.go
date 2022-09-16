package config

import (
	"log"

	"github.com/ranggabudipangestu/mvc-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/alterra-gmc?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.User{})
}
