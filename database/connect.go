package database

import (
	"go-lang-react/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:root@123@/goreact"), &gorm.Config{})
	if err != nil {
		panic("db failed to connect")
	}
	DB = connection
	connection.AutoMigrate(&models.User{}, &models.PasswordReset{})
}
