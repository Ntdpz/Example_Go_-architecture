package routers

import (
	"github.com/gofiber/fiber/v2" // นำเข้า Fiber

	"Example_Go_architecture/database"          // นำเข้าแพ็กเกจ database
	"Example_Go_architecture/internal/handlers" // นำเข้าแพ็กเกจ handlers
)

// Function SetupRoutes() ใช้สำหรับกำหนด Routes ของ API
func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler) {

	// เมื่อเข้า "/" จะแสดงข้อความ "Hello Gofiber!"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Gofiber!")
	})

	// Route สำหรับทดสอบการเชื่อมต่อฐานข้อมูล
	app.Get("/testdb", func(c *fiber.Ctx) error {
		// ดึง ข้อมูล ของ Database จาก GORM
		sqlDB, err := database.DB.DB()
		if err != nil {
			// หากเกิดข้อผิดพลาดในการดึง Database Connection ให้ส่งสถานะ 500 (Internal Server Error)
			return c.Status(500).SendString("Database connection failed")
		}

		// ทดสอบ Ping Database เพื่อตรวจสอบว่ายังเชื่อมต่ออยู่หรือไม่
		if err := sqlDB.Ping(); err != nil {
			return c.Status(500).SendString("Database connection failed")
		}

		// หากเชื่อมต่อสำเร็จ ส่งข้อความแจ้งเตือน
		return c.SendString("Successfully connected to the database!")
	})

	// กำหนด Part API

	//  API ที่เกี่ยวข้องกับ User ใช้ userHandler

	// ดึงข้อมูลผู้ใช้ทั้งหมด
	app.Get("/users", userHandler.GetUserHandler)

	// อัปเดตข้อมูลผู้ใช้
	app.Put("/users", userHandler.UpdateUserHandler)

	// ลบข้อมูลผู้ใช้
	app.Delete("/users", userHandler.DeleteUserHandler)

	// สร้างผู้ใช้ใหม่
	app.Post("/users", userHandler.CreateUserHandler)
}
