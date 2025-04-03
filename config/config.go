package config

import (
	"Example_Go_architecture/database"              // นำเข้าแพ็กเกจ database
	"Example_Go_architecture/internal/handlers"     // นำเข้าแพ็กเกจ handlers
	"Example_Go_architecture/internal/repositories" // นำเข้าแพ็กเกจ repositories
	"Example_Go_architecture/internal/services"     // นำเข้าแพ็กเกจ services
)

// ฟังค์ชั่น InitializeServices() เชื่อมโยงส่วนต่าง ๆ ของระบบ
func InitializeServices() *handlers.UserHandler {

	// ประกาศ UserRepository เพื่อเชื่อมต่อกับฐานข้อมูล
	userRepository := repositories.NewUserRepository(database.DB)

	// ประกาศ UserService โดยรับ UserRepository เพื่อใช้เรียกข้อมูล
	userService := services.NewUserService(userRepository)

	// ประกาศ UserHandler โดยรับ UserService เพื่อใช้ใน HTTP Handlers
	userHandler := handlers.NewUserHandler(userService)

	// return ค่า UserHandler เพื่อเอาไปใช้ในส่วน Router
	return userHandler
}
