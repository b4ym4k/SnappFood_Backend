package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"root/database"
	"root/routes"
)

func main() {
	//db
	database.Connect()

	app := fiber.New()

	//routes
	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
