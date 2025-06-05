package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	User_id       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"user_id" binding:"required"`
	User_name     string    `gorm:"type:varchar(250);not null" json:"user_name" binding:"required"`
	User_email    string    `gorm:"type:varchar(100);not null" json:"user_email" binding:"required"`
	User_password string    `gorm:"type:varchar(250);not null" json:"user_password" binding:"required"`

	Created_at time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp(0)" json:"updated_at"`

	// Relações
	// Categories   []Category    `gorm:"foreignKey:UserID"`
	// Transactions []Transaction `gorm:"foreignKey:UserID"`
	// Goals        []Goal        `gorm:"foreignKey:UserID"`
}
