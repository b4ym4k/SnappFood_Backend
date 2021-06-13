package controllers

import "github.com/gofiber/fiber/v2"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	//err:=c.BodyParser(&data)
	//if err != nil {
	//	return err
	//}
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	return c.JSON(data)
}
