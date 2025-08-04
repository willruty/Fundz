package finance

import "time"

type Transaction struct {
	Account_id string  `gorm:"type:uuid; primaryKey" json:"account_id"`
	Account    Account `gorm:"foreignKey:Account_id; references:Account_id" json:"user_account"`

	Category_id string   `gorm:"type:uuid; primaryKey" json:"category_id"`
	Category    Category `gorm:"foreignKey:Category_id; references:Category_id" json:"category"`

	Transaction_id string    `gorm:"type:varchar(50);primaryKey" json:"transaction_id"`
	Amount         float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	Description    string    `gorm:"type:varchar(255)" json:"description"`
	Receiver       string    `gorm:"type:varchar(100)" json:"receiver"`
	Type           string    `gorm:"type:varchar(50);not null" json:"type"`
	Date           time.Time `gorm:"type:timestamp" json:"date"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
