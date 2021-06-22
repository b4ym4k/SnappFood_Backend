package controllers

import (
	"github.com/gofiber/fiber/v2"
	"root/database"
	"root/models"
)

func ManagerCreateFood(c *fiber.Ctx) error {
	var food models.Food
	var rest	models.Restaurant
	if err := c.BodyParser(&food); err != nil {
		return err
	}

	database.DB.First(&rest, 1)

	rest.Food = append(rest.Food, food)

	database.DB.Create(&food)
	
	//restaurant := database.DB.Where("phone_number = ?", data["phoneNumber"]).First(&user)
	//database.DB.Where("id = ?", 1).Model(&user).Update("name", user.Name)

	return c.JSON(rest)
}

func ManagerDeleteFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func ManagerUpdateFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func ManagerComment(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}
