package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Goal struct {
	Category_id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"category_id" binding:"required"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	Target_amount decimal.Decimal `gorm:"type:numeric(12,2)" json:"target_amount"`
	Month_year    string          `gorm:"type:varchar(6)" json:"month_year"`
}
