package db

import (
	"github.com/akufarish/gin-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/go_web"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&models.User{}, &models.Post{})

	DB = database
}