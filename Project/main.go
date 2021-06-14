package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"root/database"
	"root/routes"
)

func main() {
	//db
	database.Connect()

	app := fiber.New()

	//ignore differance between front/back end port
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	//routes
	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
