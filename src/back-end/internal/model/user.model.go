package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"user_id"`
	UserName     string    `gorm:"type:varchar(50); not null" json:"user_name"`
	UserEmail    string    `gorm:"type:varchar(50); not null" json:"user_email"`
	UserPassword string    `gorm:"type:varchar(250); not null" json:"user_password"`
	CreatedAt    time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp(0)" json:"updated_at"`

	// Relações
	Categories   []Category    `gorm:"foreignKey:UserID"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Goals        []Goal        `gorm:"foreignKey:UserID"`
}
