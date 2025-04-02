package test

import (
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Example_Go_architecture/internal/repositories"
	"Example_Go_architecture/models"
)

func TestGetAllUsers(t *testing.T) {
	// ใช้ DSN สำหรับเชื่อมต่อกับ PostgreSQL
	dsn := "host=localhost user=admin password=admin dbname=gofiber_db port=5432 sslmode=disable"

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("ไม่สามารถเชื่อมต่อกับฐานข้อมูล: %v", err)
	}

	// ล้างข้อมูลทั้งหมดในตาราง users ก่อนเริ่มการทดสอบ
	db.Exec("DELETE FROM users")

	// สร้างตารางในฐานข้อมูลหากยังไม่มี
	db.AutoMigrate(&models.Users{})

	// สร้าง instance ของ repository
	userRepo := repositories.NewUserRepository(db)

	// เพิ่มข้อมูลตัวอย่าง
	user := models.Users{
		Username:  "john_doe",
		Password:  "secure_password123",
		Token:     "token1",
		Image:     "http://example.com/image1.jpg",
		CreatedAt: time.Now(),
	}
	db.Create(&user)

	// ทดสอบการดึงข้อมูล
	users, err := userRepo.GetAllUsers()
	if err != nil {
		t.Fatalf("เกิดข้อผิดพลาดในการดึงข้อมูลผู้ใช้: %v", err)
	}

	// ตรวจสอบจำนวนผู้ใช้ที่ได้รับว่ามีเพียง 1 รายการ
	if len(users) != 1 {
		t.Errorf("คาดหวังจำนวนผู้ใช้เป็น 1 แต่ได้ %d", len(users))
	}

	// ตรวจสอบข้อมูลของผู้ใช้
	if users[0].Username != "john_doe" {
		t.Errorf("คาดหวังชื่อผู้ใช้เป็น 'john_doe' แต่ได้ %s", users[0].Username)
	}
}
