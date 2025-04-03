package main

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/config"   // นำเข้าแพ็กเกจ config
	"Example_Go_architecture/database" // นำเข้าแพ็กเกจ database
	"Example_Go_architecture/routers"  // นำเข้าแพ็กเกจ routers
)

func main() {

	// เรียกใช้ ConnectDB() เชื่อมต่อกับ Database
	database.ConnectDB()

	// สร้าง instance ของ GoFiber เพื่อนำมาใช้เป็น Web Server
	app := fiber.New()

	// เรียกใช้งาน config
	Handler := config.InitializeServices()

	//เรียกใช้งาน Routes
	routers.SetupRoutes(app, Handler)

	// Run Server ที่ port 8080
	app.Listen(":8080")
}
