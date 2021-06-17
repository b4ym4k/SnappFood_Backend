package routes

import (
	"github.com/gofiber/fiber/v2"
	"root/controllers"
)

func Setup(app *fiber.App) {
	//manager
	app.Post("/api/manager/register", controllers.ManagerRegister)
	app.Post("/api/manager/login", controllers.ManagerLogin)
	app.Get("/api/manager/user", controllers.ManagerUser)
	app.Post("/api/manager/logout", controllers.ManagerLogout)

	//user
	app.Post("/api/user/register", controllers.UserRegister)
	app.Post("/api/user/login", controllers.UserLogin)
	app.Get("/api/user/user", controllers.UserUser)
	app.Post("/api/user/logout", controllers.UserLogout)
}
