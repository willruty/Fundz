package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Goal struct {
	GoalID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"goal_id"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CategoryID uuid.UUID `gorm:"type:uuid" json:"category_id"`
	Category   Category  `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`

	TargetAmount decimal.Decimal `gorm:"type:numeric(12,2)" json:"target_amount"`
	MonthYear    string          `gorm:"type:varchar(6)" json:"month_year"`
}
