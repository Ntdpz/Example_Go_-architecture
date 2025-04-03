package repositories

import (
	"time"

	"gorm.io/gorm"

	"Example_Go_architecture/models" // นำเข้า models
)

// สร้าง interface สำหรับ UserRepository เพื่อเอาไว้เรียกใช้
type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	FindUserByParams(username, token string) ([]models.Users, error)
	UpdateUserByToken(token string, user models.Users) (*models.Users, error)
	DeleteUserByToken(token string) error
	CreateUser(user models.Users) (*models.Users, error)
}

// สร้าง struct สำหรับจัดการข้อมูล
type userRepository struct {
	DB *gorm.DB // กำหนด Gorm DB เชื่อมต่อ Database
}

// ฟังก์ชันสร้าง UserRepository ใหม่
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

// ฟังก์ชันดึงข้อมูล User ทั้งหมด
func (r *userRepository) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err // ถ้ามีข้อผิดพลาดให้คืนค่าผลลัพธ์เป็น error
	}
	return users, nil // คืนค่าผลลัพธ์เป็น User ทั้งหมด
}

// ฟังก์ชันค้นหาผู้ใช้ตาม Paramiter (Username หรือ Token)
func (r *userRepository) FindUserByParams(username, token string) ([]models.Users, error) {
	var users []models.Users
	query := r.DB.Model(&models.Users{}) // สร้าง query สำหรับค้นหาผู้ใช้

	// ถ้ามีการส่ง Username เข้ามา
	if username != "" {
		query = query.Where("username = ?", username)
	}
	// ถ้ามีการส่ง Token เข้ามา
	if token != "" {
		query = query.Where("token = ?", token)
	}

	// ดึงข้อมูลผู้ใช้ที่ตรงกับเงื่อนไข
	if err := query.Find(&users).Error; err != nil {
		return nil, err // ถ้ามีข้อผิดพลาดให้คืนค่าผลลัพธ์เป็น error
	}
	return users, nil // คืนค่าผลลัพธ์เป็นผู้ใช้ที่ตรงกับเงื่อนไข
}

// ฟังก์ชันอัปเดตข้อมูล User ตาม Token
func (r *userRepository) UpdateUserByToken(token string, user models.Users) (*models.Users, error) {
	var existingUser models.Users

	// ค้นหาผู้ใช้ตาม Token
	if err := r.DB.Where("token = ?", token).First(&existingUser).Error; err != nil {
		return nil, err // ถ้าผู้ใช้ไม่พบให้คืนค่าผลลัพธ์เป็น error
	}

	// ตรวจสอบข้อมูลที่ส่งมา ถ้าไม่มีข้อมูลให้ใช้ข้อมูลเดิม
	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.Image != "" {
		existingUser.Image = user.Image
	}

	// อัปเดตเวลาที่มีการเปลี่ยนแปลง
	existingUser.UpdatedAt = time.Now()

	// บันทึกการเปลี่ยนแปลง
	if err := r.DB.Save(&existingUser).Error; err != nil {
		return nil, err // ถ้าการบันทึกไม่สำเร็จให้คืนค่าผลลัพธ์เป็น error
	}

	return &existingUser, nil // คืนค่าผลลัพธ์เป็นผู้ใช้ที่ถูกอัปเดต
}

// ฟังก์ชันลบ User ตาม Token
func (r *userRepository) DeleteUserByToken(token string) error {
	var user models.Users
	// ค้นหาผู้ใช้ตาม Token
	if err := r.DB.Where("token = ?", token).First(&user).Error; err != nil {
		return err // ถ้าผู้ใช้ไม่พบให้คืนค่าผลลัพธ์เป็น error
	}

	// ลบผู้ใช้จากฐานข้อมูล
	if err := r.DB.Delete(&user).Error; err != nil {
		return err // ถ้าการลบไม่สำเร็จให้คืนค่าผลลัพธ์เป็น error
	}

	return nil // คืนค่า nil ถ้าลบสำเร็จ
}

// ฟังก์ชันสร้าง User ใหม่
func (r *userRepository) CreateUser(user models.Users) (*models.Users, error) {
	// บันทึกข้อมูลผู้ใช้ใหม่ในฐานข้อมูล
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err // ถ้าการบันทึกไม่สำเร็จให้คืนค่าผลลัพธ์เป็น error
	}
	return &user, nil // คืนค่าผลลัพธ์เป็นผู้ใช้ใหม่ที่ถูกสร้าง
}
