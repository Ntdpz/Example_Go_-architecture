package routers

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/database"
)

func SetupRoutes(app *fiber.App) {
	// กำหนด route หลัก
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Gofiber!")
	})

	// กำหนด route สำหรับทดสอบการเชื่อมต่อฐานข้อมูล
	app.Get("/testdb", func(c *fiber.Ctx) error {
		err := database.DB.DB().Ping()
		if err != nil {
			return c.Status(500).SendString("Database connection failed")
		}
		return c.SendString("Successfully connected to the database!")
	})
}
