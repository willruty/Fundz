package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	Transaction_id          uuid.UUID       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"transaction_id"`
	User_id                 uuid.UUID       `gorm:"type:uuid; foreingKey" json:"user_id"`
	Category_id             uuid.UUID       `gorm:"type:uuid; foreingKey" json:"category_id"`
	Transaction_type        string          `gorm:"type:varchar(7); not null" json:"transaction_type"`
	Transaction_amount      decimal.Decimal `gorm:"type:numeric(12,4)" json:"transaction_amount"`
	Transaction_description string          `gorm:"type:varchar(100); not null" json:"transaction_description"`
	Transaction_date        time.Time       `gorm:"type:timestamp(0)" json:"transaction_date"`
}
