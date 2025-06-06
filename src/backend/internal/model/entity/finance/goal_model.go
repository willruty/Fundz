package entity

import (
	"fundz/internal/model/entity"

	"github.com/shopspring/decimal"
)

type Goal struct {
	Category_id string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"category_id" binding:"required"`

	UserID string      `gorm:"type:uuid" json:"user_id"`
	User   entity.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CategoryID string   `gorm:"type:uuid" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	Target_amount decimal.Decimal `gorm:"type:numeric(12,2)" json:"target_amount"`
	Month_year    string          `gorm:"type:varchar(6)" json:"month_year"`
}
