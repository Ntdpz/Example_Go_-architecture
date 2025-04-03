package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres" // ใช้ GORM กับ PostgreSQL เป็น Database
	"gorm.io/gorm"            // นำเข้า GORM
)

// ประกาศตัวแปร DB เพื่อให้สามารถใช้เชื่อมต่อDatabaseได้จากทุกที่ในโปรเจกต์
var DB *gorm.DB

// ConnectDB() เป็น Function สำหรับเชื่อมต่อกับ Database
func ConnectDB() {

	// กำหนดข้อมูลการเชื่อมต่อ Database
	dsn := "host=psql_db user=admin password=admin dbname=gofiber_db port=5432 sslmode=disable"

	var err error

	// เชื่อมต่อ Database โดยใช้ GORM และ PostgreSQL driver
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// ถ้าเชื่อมต่อไม่สำเร็จ ให้แสดงข้อความ error และหยุดโปรแกรม
		log.Fatalf("Error connecting to the database: %v", err)
		return
	}

	// หากเชื่อมต่อสำเร็จ ให้แสดงข้อความแจ้งเตือน
	fmt.Println("Successfully connected to the database!")
}
