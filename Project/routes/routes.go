package routes

import (
	"github.com/gofiber/fiber/v2"
	"root/controllers"
)

func Setup(app *fiber.App) {
	//==================Manager================================================

	//authController
	app.Post("/api/manager/register", controllers.ManagerRegister)
	app.Post("/api/manager/login", controllers.ManagerLogin)
	app.Get("/api/manager/user", controllers.ManagerUser)
	app.Post("/api/manager/logout", controllers.ManagerLogout)
	app.Put("/api/manager/updateProfile", controllers.ManagerUpdateProfile) //not implemented

	//managerController
	app.Post("/api/manager/createFood", controllers.ManagerCreateFood)   //not implemented
	app.Delete("/api/manager/deleteFood", controllers.ManagerDeleteFood) //not implemented
	app.Delete("/api/manager/UpdateFood", controllers.ManagerUpdateFood) //not implemented
	app.Post("/api/manager/comment", controllers.ManagerCommect)         //not implemented

	//==================User==================================================

	//authController
	app.Post("/api/user/register", controllers.UserRegister)
	app.Post("/api/user/login", controllers.UserLogin)
	app.Get("/api/user/user", controllers.UserUser)
	app.Post("/api/user/logout", controllers.UserLogout)
	app.Put("/api/manager/updateProfile", controllers.UserUpdateProfile) //not implemented

}
