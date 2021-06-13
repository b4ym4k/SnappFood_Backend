package routes

import (
	"github.com/gofiber/fiber/v2"
	"root/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
