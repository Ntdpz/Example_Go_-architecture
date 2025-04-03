package handlers

import (
	"github.com/gofiber/fiber/v2"

	"Example_Go_architecture/internal/services" // นำเข้า service
	"Example_Go_architecture/middlewares"       // นำเข้า middleware
	"Example_Go_architecture/models"            // นำเข้า model
)

// สร้าง struct UserHandler เก็บ service
type UserHandler struct {
	Service services.UserService // กำหนด field สำหรับเรียกใช้งาน UserService
}

// ฟังก์ชันสร้าง UserHandler ใหม่
func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service} // คืนค่า UserHandler ที่สร้างขึ้นใหม่
}

// ฟังก์ชันดึงข้อมูลผู้ใช้
func (h *UserHandler) GetUserHandler(c *fiber.Ctx) error {
	// รับค่า parameter Username และ Token
	username := c.Query("Username")
	token := c.Query("Token")

	// ถ้ามีการส่ง Username หรือ Token
	if username != "" || token != "" {
		// เรียก Service เพื่อทำงานตาม Service
		users, err := h.Service.GetUsersByParams(username, token)
		if err != nil {
			// ถ้ามีข้อผิดพลาดให้คืนค่า 500 (Internal Server Error)
			return c.Status(500).JSON(fiber.Map{
				"error": "Failed to retrieve user",
			})
		}
		// ถ้าผู้ใช้ไม่พบ
		if len(users) == 0 {
			return c.Status(404).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		// คืนค่าผลลัพธ์
		return c.JSON(fiber.Map{
			"count": len(users), // จำนวนผู้ใช้
			"data":  users,      // ข้อมูลผู้ใช้
		})
	}

	// ถ้าไม่ที parameter จะดึงข้อมูลผู้ใช้ทั้งหมด
	users, err := h.Service.GetUsers()
	if err != nil {
		// ถ้ามีข้อผิดพลาดให้คืนค่า 500 (Internal Server Error)
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to retrieve users",
		})
	}
	// คืนค่าผลลัพธ์
	return c.JSON(fiber.Map{
		"count": len(users), // จำนวนผู้ใช้ทั้งหมด
		"data":  users,      // ข้อมูลผู้ใช้ทั้งหมด
	})
}

// ฟังก์ชัน Update User
func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	// รับค่า Token จาก parameter
	token := c.Query("Token")

	// ถ้าไม่มีการส่ง Token มาให้คืนค่าผลลัพธ์เป็น 400 (Bad Request)
	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	var userData models.Users

	// ดึงข้อมูลจาก Body (JSON)
	if err := c.BodyParser(&userData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// อัปเดตข้อมูลผู้ใช้โดยใช้ Token (Service)
	updatedUser, err := h.Service.UpdateUserByToken(token, userData)
	if err != nil {
		// ถ้าผู้ใช้ไม่พบให้คืนค่า 404 (Not Found)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// return ผลลัพธ์เมื่ออัปเดตสำเร็จ
	return c.JSON(fiber.Map{
		"data":    updatedUser,                 // ข้อมูลผู้ใช้ที่ถูกอัปเดต
		"message": "User updated successfully", // ข้อความสำเร็จ
	})
}

// ฟังก์ชัน Delete User
func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	// รับค่า Token จาก parameter
	token := c.Query("Token")

	// ถ้าไม่มีการส่ง Token มาให้คืนค่า 400 (Bad Request)
	if token == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Token is required",
		})
	}

	// เรียก Service เพื่อลบข้อมูลผู้ใช้
	err := h.Service.DeleteUserByToken(token)
	if err != nil {
		// ถ้าผู้ใช้ไม่พบให้คืนค่า 404 (Not Found)
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//  return ลัพธ์เมื่อลบสำเร็จ
	return c.JSON(fiber.Map{
		"message": "User deleted successfully", // ข้อความสำเร็จ
	})
}

// ฟังก์ชัน Create User
func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var userData models.Users

	// รับ JSON จาก Body
	if err := c.BodyParser(&userData); err != nil {
		// ถ้าข้อมูลใน Body ผิดพลาดให้คืนค่า 400 (Bad Request)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// ตรวจสอบว่า Username และ Password ถูกส่งมาหรือไม่
	if userData.Username == "" || userData.Password == "" {
		// ถ้าไม่ครบให้คืนค่า 400 (Bad Request)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Username and Password are required",
		})
	}

	// เรียก Middleware เพื่อสร้าง Token
	token, err := middlewares.GenerateToken(userData.Username)
	if err != nil {
		// ถ้ามีข้อผิดพลาดในการสร้าง Token ให้คืนค่า 500 (Internal Server Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}
	userData.Token = token

	// บันทึกข้อมูลผู้ใช้ใหม่ลงใน Database
	newUser, err := h.Service.CreateUser(userData)

	if err != nil {
		// ถ้ามีข้อผิดพลาดในการบันทึกข้อมูลให้คืนค่า 500 (Internal Server Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	// คืนค่าผลลัพธ์เมื่อสร้างผู้ใช้สำเร็จ
	return c.JSON(fiber.Map{
		"data":    newUser,                     // ข้อมูลผู้ใช้ที่ถูกสร้าง
		"message": "User created successfully", // ข้อความสำเร็จ
	})
}
