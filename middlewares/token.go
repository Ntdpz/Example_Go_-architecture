package middlewares

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"

	"Example_Go_architecture/database"
	"Example_Go_architecture/models"
)

// กำหนดค่า Secret Key จากไฟล์ .env
var secretKey string

// ฟังก์ชันสำหรับโหลดค่า .env
func init() {
	// โหลดไฟล์ .env
	err := godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// อ่านค่า SECRET_KEY จากไฟล์ .env
	secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("SECRET_KEY is required in .env file")
	}
}

// Function GenerateToken() ใช้สำหรับสร้าง JWT Token
func GenerateToken(username string) (string, error) {
	// สุ่มตัวเลข 6 หลัก
	rand.Seed(time.Now().UnixNano())                     // ตั้งค่า Seed เพื่อให้เลขที่สุ่มแตกต่างกันทุกครั้ง
	randomNum := fmt.Sprintf("%06d", rand.Intn(1000000)) // สร้างเลขสุ่ม 6 หลัก

	// แบ่ง Username ออกเป็น 2 ส่วน
	mid := len(username) / 2
	part1 := username[:mid]
	part2 := username[mid:]

	// สร้าง Token แบบดิบ (Raw Token)
	rawToken := fmt.Sprintf("%s%s%s%s", randomNum, part1, secretKey, part2)

	// สร้าง JWT Token พร้อมกำหนด Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"raw": rawToken,                              // เก็บค่า Token ที่สร้างขึ้น
		"exp": time.Now().Add(time.Hour * 24).Unix(), // หมดอายุใน 24 ชั่วโมง
	})

	// เซ็ต Token ด้วย Secret Key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err // หากเกิดข้อผิดพลาดให้คืนค่า Error
	}

	return tokenString, nil // คืนค่า JWT Token
}

// CheckTokenMiddleware คือ Middleware ที่ใช้ตรวจสอบว่า Token มีใน Header และถูกต้องหรือไม่
func CheckTokenMiddleware(c *fiber.Ctx) error {
	// ดึงค่า Token จาก Header
	token := c.Get("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token is required",
		})
	}

	// ลบคำว่า "Bearer " ออกจาก Token หากมี
	if len(token) > 6 && strings.ToUpper(token[:7]) == "BEARER " {
		token = token[7:]
	}

	// ตรวจสอบว่า Token นี้มีใน Database หรือไม่
	var user models.Users
	if err := database.DB.Where("token = ?", token).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// ถ้า Token ถูกต้องให้ทำการส่งต่อไปยัง Handler
	return c.Next()
}
