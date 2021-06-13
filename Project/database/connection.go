package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	//db
	_, err := gorm.Open(mysql.Open("root:rootroot@/snapfood_go_auth"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
