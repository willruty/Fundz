package finance

import "time"

type Account struct {
	Account_id string `gorm:"type:varchar(50);primaryKey" json:"account_id"`
	// user id
	Account_name string `gorm:"type:varchar(50);not null" json:"account_name"`
	Account_type string `gorm:"type:varchar(50);not null" json:"account_type"` // "corrente", "poupança"
	Bank_name    string `gorm:"type:varchar(50);not null" json:"bank_name"`    // "Banco do Brasil", "Caixa Econômica Federal", etc.

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
