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
	app.Put("/api/manager/updateFood", controllers.ManagerUpdateFood)    //not implemented
	app.Post("/api/manager/comment", controllers.ManagerComment)         //not implemented

	//managerOrderController
	app.Put("/api/manager/checkOrder", controllers.ManagerCheckOrder)                        //not implemented
	app.Get("/api/manager/findNameAndDistrictById/:id", controllers.FindNameAndDistrictById) //not implemented
	app.Get("/api/manager/orders/:id", controllers.GetOrders)                                //not implemented
	app.Get("/api/manager/foods/:id", controllers.GetManagerFoods)                           //not implemented
	app.Get("/api/manager/comment/:id", controllers.GetCommects)                             //not implemented

	//==================User==================================================

	//authController
	app.Post("/api/user/register", controllers.UserRegister)
	app.Post("/api/user/login", controllers.UserLogin)
	app.Get("/api/user/user", controllers.UserUser)
	app.Post("/api/user/logout", controllers.UserLogout)
	app.Put("/api/manager/updateProfile", controllers.UserUpdateProfile) //not implemented

	//userController
	app.Post("/api/user/allFood", controllers.UserGetAllFood)     //not implemented
	app.Get("/api/user/getOrders/:id", controllers.UserGetOrders) //not implemented
	app.Post("/api/user/order", controllers.UserCreateOrder)      //not implemented
	app.Post("/api/user/history", controllers.UserGetHistory)     //not implemented
	app.Post("/api/user/favorites", controllers.UserGetFavorites) //not implemented
	app.Post("/api/user/comment", controllers.UserCreateComment)  //not implemented

}
