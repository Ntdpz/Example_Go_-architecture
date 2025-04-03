package services

import (
	"Example_Go_architecture/internal/repositories" // นำเข้า repository
	"Example_Go_architecture/models"                // นำเข้า models
)

// สร้าง interface สำหรับ UserService สำหรับนำไปเรียกใช้
type UserService interface {
	GetUsers() ([]models.Users, error)
	GetUsersByParams(username, token string) ([]models.Users, error)
	UpdateUserByToken(token string, user models.Users) (*models.Users, error)
	DeleteUserByToken(token string) error
	CreateUser(user models.Users) (*models.Users, error)
}

// สร้าง struct สำหรับ UserService โดยใช้ repository สำหรับจัดการข้อมูล
type userService struct {
	repo repositories.UserRepository
}

// ฟังก์ชันสร้าง UserService ใหม่
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// ฟังก์ชันดึงข้อมูลผู้ใช้ทั้งหมด
func (s *userService) GetUsers() ([]models.Users, error) {
	return s.repo.GetAllUsers() // เรียกใช้ repository เพื่อดึงข้อมูลผู้ใช้ทั้งหมด
}

// ฟังก์ชันค้นหาผู้ใช้ตามพารามิเตอร์ (Username หรือ Token)
func (s *userService) GetUsersByParams(username, token string) ([]models.Users, error) {
	return s.repo.FindUserByParams(username, token) // เรียกใช้ repository เพื่อค้นหาผู้ใช้ตามพารามิเตอร์
}

// ฟังก์ชันอัปเดตข้อมูลผู้ใช้ตาม Token
func (s *userService) UpdateUserByToken(token string, user models.Users) (*models.Users, error) {
	return s.repo.UpdateUserByToken(token, user) // เรียกใช้ repository เพื่ออัปเดตข้อมูลผู้ใช้ตาม Token
}

// ฟังก์ชันลบผู้ใช้ตาม Token
func (s *userService) DeleteUserByToken(token string) error {
	return s.repo.DeleteUserByToken(token) // เรียกใช้ repository เพื่อทำการลบผู้ใช้ตาม Token
}

// ฟังก์ชันสร้างผู้ใช้ใหม่
func (s *userService) CreateUser(user models.Users) (*models.Users, error) {
	return s.repo.CreateUser(user) // เรียกใช้ repository เพื่อสร้างผู้ใช้ใหม่ในฐานข้อมูล
}
