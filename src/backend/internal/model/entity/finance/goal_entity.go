package finance

import (
	"time"

	"github.com/shopspring/decimal"
)

type Goal struct {
	GoalId        string          `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"goal_id"`
	AccountId     string          `gorm:"type:uuid; primaryKey" json:"account_id"`
	CategoryId    string          `gorm:"type:uuid; primaryKey" json:"category_id"`
	TargetAmount  decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"target_amount"`
	CurrentAmount decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"current_amount"`
	GoalTitle     string          `gorm:"type:varchar(100);not null" json:"goal_title"`
	Description   string          `gorm:"type:varchar(255)" json:"description"`
	Deadline      time.Time       `gorm:"type:timestamp" json:"deadline"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
