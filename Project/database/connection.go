package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"root/models"
)

func Connect() {
	//db
	db, err := gorm.Open(mysql.Open("root:rootroot@/snapfood_go_auth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
}
