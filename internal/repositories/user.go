package repositories

import (
	"gorm.io/gorm"

	"Example_Go_architecture/models"
)

type UserRepository interface {
	GetAllUsers() ([]models.Users, error)
	FindUserByParams(username, token string) ([]models.Users, error)
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
