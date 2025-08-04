package finance

import (
	usuario "fundz/internal/model/entity/user"
	"time"
)

type Account struct {
	Account_id string `gorm:"type:varchar(50);primaryKey" json:"account_id"`

	User_id     string              `gorm:"type:uuid; primaryKey" json:"user_id"`
	UserAccount usuario.UserAccount `gorm:"foreignKey:User_id;references:User_id" json:"user_account"`

	Account_name string `gorm:"type:varchar(50);not null" json:"account_name"`
	Account_type string `gorm:"type:varchar(50);not null" json:"account_type"`
	Bank_name    string `gorm:"type:varchar(50);not null" json:"bank_name"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
