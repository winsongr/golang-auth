package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:root@123/goreact"), &gorm.Config{})
	if err != nil {
		panic("db failed to connect")
	}
}
