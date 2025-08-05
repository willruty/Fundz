package finance

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionId string          `gorm:"type:varchar(50);primaryKey" json:"transaction_id"`
	ActId     string          `gorm:"type:text" json:"account_id"`
	CategoryId    string          `gorm:"type:text" json:"category_id"`
	Amount        decimal.Decimal `gorm:"type:decimal(10,2);not null" json:"amount"`
	Description   string          `gorm:"type:varchar(255)" json:"description"`
	Receiver      string          `gorm:"type:varchar(100)" json:"receiver"`
	Type          string          `gorm:"type:varchar(50);not null" json:"type"`
	Date          time.Time       `gorm:"type:timestamp" json:"date"`

	Created_at    time.Time       `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at    time.Time       `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
