package entity

import (
	"fundz/internal/model/entity"

	"github.com/google/uuid"
)

type Category struct {
	Category_id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"category_id" binding:"required"`
	User_id       uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User          entity.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Category_name string    `gorm:"type:varchar(100)" json:"category_name"`

	// Transactions []Transaction `gorm:"foreignKey:CategoryID"`
	// Goals        []Goal        `gorm:"foreignKey:CategoryID"`
}
