package routers

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/database"
	"Example_Go_architecture/internal/handlers"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Gofiber!")
	})

	app.Get("/testdb", func(c *fiber.Ctx) error {
		sqlDB, err := database.DB.DB()
		if err != nil {
			return c.Status(500).SendString("Database connection failed")
		}

		if err := sqlDB.Ping(); err != nil {
			return c.Status(500).SendString("Database connection failed")
		}

		return c.SendString("Successfully connected to the database!")
	})

	app.Get("/user", userHandler.GetUserHandler)

}
