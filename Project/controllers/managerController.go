package controllers

import (
	"github.com/gofiber/fiber/v2"
	"root/models"
)

func ManagerCreateFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
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
