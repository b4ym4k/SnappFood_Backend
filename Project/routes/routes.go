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
	app.Put("/api/manager/updateProfile", controllers.ManagerUpdateProfile)

	//managerController
	app.Post("/api/manager/createFood", controllers.ManagerCreateFood)
	app.Delete("/api/manager/deleteFood", controllers.ManagerDeleteFood)
	app.Put("/api/manager/updateFood", controllers.ManagerUpdateFood)
	app.Post("/api/manager/comment", controllers.ManagerComment)

	//managerOrderController
	app.Put("/api/manager/checkOrder", controllers.ManagerCheckOrder)
	app.Get("/api/manager/findNameAndDistrictById/:id", controllers.FindNameAndDistrictById)
	app.Get("/api/manager/orders/:id", controllers.GetOrders)
	app.Get("/api/manager/foods/:id", controllers.GetManagerFoods)
	app.Get("/api/manager/comment/:id", controllers.GetCommects)

	//==================User==================================================

	//authController
	app.Post("/api/user/register", controllers.UserRegister)
	app.Post("/api/user/login", controllers.UserLogin)
	app.Get("/api/user/user", controllers.UserUser)
	app.Post("/api/user/logout", controllers.UserLogout)
	app.Put("/api/user/updateProfile", controllers.UserUpdateProfile)

	//userController
	app.Post("/api/user/allFood", controllers.UserGetAllFood)
	app.Get("/api/user/getOrders/:id", controllers.UserGetOrders)
	app.Post("/api/user/order", controllers.UserCreateOrder)
	app.Post("/api/user/history", controllers.UserGetHistory)
	app.Post("/api/user/favorites", controllers.UserGetFavorites)
	app.Post("/api/user/comment", controllers.UserCreateComment)

}
