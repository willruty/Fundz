package finance

import "time"

type Transaction struct {
	// account id
	// category id
	Transaction_id string `gorm:"type:varchar(50);primaryKey" json:"transaction_id"`
	Amount		float64 `gorm:"type:decimal(10,2);not null" json:"amount"` // positive for income, negative for expense
	Description string `gorm:"type:varchar(255)" json:"description"` // description of the transaction
	Receiver string `gorm:"type:varchar(100)" json:"receiver"` // receiver of the transaction, can be a person or a company
	Type string `gorm:"type:varchar(50);not null" json:"type"` // "income", "expense", "transfer"
	Date time.Time `gorm:"type:timestamp" json:"date"` // date of the transaction

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}

