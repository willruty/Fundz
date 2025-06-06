package entity

import (
	"fundz/internal/model/entity"
)

type Category struct {
	Category_id   string      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"category_id" binding:"required"`
	User_id       string      `gorm:"type:uuid" json:"user_id"`
	User          entity.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Category_name string      `gorm:"type:varchar(100)" json:"category_name"`

}
