package models

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Token     string    `gorm:"type:varchar(255);" json:"token,omitempty"`
	Image     string    `gorm:"type:text;" json:"image,omitempty"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
