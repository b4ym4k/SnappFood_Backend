package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"root/models"
)

var DB *gorm.DB

func Connect() {
	//db
	Connection, err := gorm.Open(mysql.Open("root:rootroot@/snapfood_go_auth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = Connection
	Connection.AutoMigrate(&models.Manager{})
	Connection.AutoMigrate(&models.User{})
}
