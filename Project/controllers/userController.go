package controllers

import (
	"github.com/gofiber/fiber/v2"
	"root/models"
)

func UserGetAllFood(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserGetOrders(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserCreateOrder(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserGetHistory(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserGetFavorites(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}

func UserCreateComment(c *fiber.Ctx) error {
	var user models.User
	return c.JSON(user)
}
