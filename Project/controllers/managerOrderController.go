package controllers

import (
	"github.com/gofiber/fiber/v2"
	"root/models"
)

func ManagerCheckOrder(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func FindNameAndDistrictById(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func GetOrders(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func GetManagerFoods(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func GetCommects(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}
