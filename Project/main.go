package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
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

	// Or extend your config for customization
	app.Use(favicon.New(favicon.Config{
		File: "./assets/food.png",
	}))

	//routes
	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
