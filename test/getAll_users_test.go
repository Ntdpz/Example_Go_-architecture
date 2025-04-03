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
	// ตั้งค่า DSN สำหรับเชื่อมต่อกับ PostgreSQL
	dsn := "host=localhost user=admin password=admin dbname=gofiber_db port=5432 sslmode=disable"

	// เชื่อมต่อกับ Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("ไม่สามารถเชื่อมต่อกับฐานข้อมูล: %v", err)
	}

	// ล้างข้อมูลทั้งหมดในตาราง users ก่อนเริ่ม Test
	if err := db.Exec("DELETE FROM users").Error; err != nil {
		t.Fatalf("ไม่สามารถล้างข้อมูลในตาราง users: %v", err)
	}

	// สร้างตารางในฐานข้อมูลหากยังไม่มี
	if err := db.AutoMigrate(&models.Users{}); err != nil {
		t.Fatalf("ไม่สามารถสร้างตาราง users: %v", err)
	}

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
	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("ไม่สามารถเพิ่มผู้ใช้ลงในฐานข้อมูล: %v", err)
	}

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

	// Cleanup: ลบข้อมูลหลังจากการทดสอบ
	t.Cleanup(func() {
		if err := db.Exec("DELETE FROM users").Error; err != nil {
			t.Errorf("ไม่สามารถลบข้อมูลหลังการทดสอบ: %v", err)
		}
	})
}
