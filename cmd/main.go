package main

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/database"
	"Example_Go_architecture/routers"
)

func main() {

	database.ConnectDB()

	app := fiber.New()

	routers.SetupRoutes(app)

	app.Listen(":8080")
}
