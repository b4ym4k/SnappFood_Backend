package main

import (
	"log"
	"root/database"
	"root/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//db
	database.Connect()

	app := fiber.New()

	//routes
	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
