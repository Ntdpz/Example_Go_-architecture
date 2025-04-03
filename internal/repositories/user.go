package repositories

import (
	"time"

	"gorm.io/gorm"

	"Example_Go_architecture/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	FindUserByParams(username, token string) ([]models.Users, error)
	UpdateUserByToken(token string, user models.Users) (*models.Users, error)
	DeleteUserByToken(token string) error
}
type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}
func (r *userRepository) GetAllUsers() ([]models.Users, error) {
	var users []models.Users
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindUserByParams(username, token string) ([]models.Users, error) {
	var users []models.Users
	query := r.DB.Model(&models.Users{})

	if username != "" {
		query = query.Where("username = ?", username)
	}
	if token != "" {
		query = query.Where("token = ?", token)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) UpdateUserByToken(token string, user models.Users) (*models.Users, error) {
	var existingUser models.Users

	// ค้นหาผู้ใช้ตาม Token
	if err := r.DB.Where("token = ?", token).First(&existingUser).Error; err != nil {
		return nil, err
	}

	// ตรวจสอบว่าได้ส่งข้อมูลมาอัปเดตหรือไม่ ถ้าไม่ส่งมาก็ใช้ค่าเดิม
	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.Image != "" {
		existingUser.Image = user.Image
	}

	// อัปเดต UpdatedAt ให้อัตโนมัติ
	existingUser.UpdatedAt = time.Now()

	// บันทึกการเปลี่ยนแปลง
	if err := r.DB.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

// ฟังก์ชันลบผู้ใช้ตาม Token
func (r *userRepository) DeleteUserByToken(token string) error {
	var user models.Users
	if err := r.DB.Where("token = ?", token).First(&user).Error; err != nil {
		return err
	}

	// ลบผู้ใช้
	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
