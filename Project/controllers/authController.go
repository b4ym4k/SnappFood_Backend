package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"root/database"
	"root/models"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//err:=c.BodyParser(&data)
	//if err != nil {
	//	return err
	//}
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password, //data["password"],
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
