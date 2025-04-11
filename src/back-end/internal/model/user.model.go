package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	User_id       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"user_id"`
	User_name     string    `gorm:"type:varchar(50)" json:"user_name"`
	User_email    string    `gorm:"type:varchar(50)" json:"user_email"`
	User_password string    `gorm:"type:varchar(50)" json:"user_password"`
	Created_at    time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	Update_at     time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
}
