package controllers

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World! ✌🏻")
}
