package routes

import (
	"github.com/gofiber/fiber/v2"
	"root/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/manager/register", controllers.ManagerRegister)
	app.Post("/api/manager/login", controllers.ManagerLogin)
	app.Get("/api/manager/user", controllers.ManagerUser)
	app.Post("/api/manager/logout", controllers.ManagerLogout)
}
