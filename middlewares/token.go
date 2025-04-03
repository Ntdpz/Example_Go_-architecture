package middlewares

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var secretKey = "my%ona)j8&6m" // ค่าของฉัน

func GenerateToken(username string) (string, error) {
	// สุ่มตัวเลข 6 หลัก
	rand.Seed(time.Now().UnixNano())
	randomNum := fmt.Sprintf("%06d", rand.Intn(1000000))

	// แบ่ง Username เป็นสองส่วน
	mid := len(username) / 2
	part1 := username[:mid]
	part2 := username[mid:] // ส่วนหลัง

	// รวม Token Format
	rawToken := fmt.Sprintf("%s%s%s%s", randomNum, part1, secretKey, part2)

	// สร้าง JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"raw": rawToken,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // หมดอายุใน 24 ชั่วโมง
	})

	// เซ็น Token ด้วย Secret Key (จาก ENV หรือใช้ค่าคงที่)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
