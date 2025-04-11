package model

import "github.com/google/uuid"

type Category struct {
	Category_id   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"category_id"`
	User_id       uuid.UUID `gorm:"type:uuid; foreingKey" json:"user_id"`
	Category_name string    `gorm:"type:varchar(30)" json:"category_name"`
}
