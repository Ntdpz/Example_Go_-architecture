package main

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/config"
	"Example_Go_architecture/database"
	"Example_Go_architecture/routers"

)

func main() {

	database.ConnectDB()

	app := fiber.New()

	Handler := config.InitializeServices()

	routers.SetupRoutes(app, Handler)

	app.Listen(":8080")
}
