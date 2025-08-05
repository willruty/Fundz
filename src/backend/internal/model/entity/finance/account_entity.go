package finance

import (
	"time"
)

type Account struct {
	UserId    string `gorm:"type:uuid; primaryKey" json:"user_id"`
	AccountId string `gorm:"type:varchar(50);primaryKey" json:"account_id"`
	AccountName string `gorm:"type:varchar(50);not null" json:"account_name"`
	AccountType string `gorm:"type:varchar(50);not null" json:"account_type"`
	BankName    string `gorm:"type:varchar(50);not null" json:"bank_name"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
