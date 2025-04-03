package middlewares

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt" // นำเข้า JWT Library
)

// กำหนดค่า Secret Key จากไฟล์ .env
var secretKey string

// ฟังก์ชัน init จะทำงานอัตโนมัติเมื่อเริ่มทำงาน
func init() {
	secretKey = os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		panic("JWT_SECRET_KEY is not set in the environment variables") // หยุดโปรแกรมถ้าไม่พบค่าใน ENV
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
