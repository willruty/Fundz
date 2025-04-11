package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Goal struct {
	Goal_id       uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"goal_id"`
	Category_id   uuid.UUID       `gorm:"type:uuid; foreingKey" json:"category_id"`
	User_id       uuid.UUID       `gorm:"type:uuid; foreingKey" json:"user_id"`
	Target_amount decimal.Decimal `gorm:"type:numeric(12,4)" json:"target_amount"`
	Month_year    string          `gorm:"type:varchar(6)" json:"month_year"`
}
