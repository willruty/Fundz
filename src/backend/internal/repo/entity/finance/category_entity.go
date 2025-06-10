package finance

import "time"

type Category struct {
	Category_id string `gorm:"type:varchar(50);primaryKey" json:"category_id"`

	Account_id string  `gorm:"type:uuid; primaryKey" json:"account_id"`
	Account    Account `gorm:"foreignKey:Account_id; references:Account_id" json:"user_account"`

	Category_title string `gorm:"type:varchar(50);not null" json:"category_title"`
	Description    string `gorm:"type:varchar(255)" json:"description"`
	Color_code     string `gorm:"type:varchar(7)" json:"color_code"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
