package finance

import "time"

type Transaction struct {
	TransactionId string    `gorm:"type:varchar(50);primaryKey" json:"transaction_id"`
	AccountId     string    `gorm:"type:text" json:"account_id"`
	CategoryId    string    `gorm:"type:text" json:"category_id"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Description   string    `gorm:"type:varchar(255)" json:"description"`
	Receiver      string    `gorm:"type:varchar(100)" json:"receiver"`
	Type          string    `gorm:"type:varchar(50);not null" json:"type"`
	Date          time.Time `gorm:"type:timestamp" json:"date"`
	Created_at    time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at    time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
