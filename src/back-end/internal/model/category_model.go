package model

import "github.com/google/uuid"

type Category struct {
	CategoryID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"category_id"`
	UserID       uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User         User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CategoryName string    `gorm:"type:varchar(30)" json:"category_name"`

	Transactions []Transaction `gorm:"foreignKey:CategoryID"`
	Goals        []Goal        `gorm:"foreignKey:CategoryID"`
}
