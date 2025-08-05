package finance

import "time"

type Category struct {
	CategoryId    string `gorm:"type:varchar(50);primaryKey" json:"category_id"`
	AccountId     string `gorm:"type:uuid; primaryKey" json:"account_id"`
	CategoryTitle string `gorm:"type:varchar(50);not null" json:"category_title"`
	Description   string `gorm:"type:varchar(255)" json:"description"`
	ColorCode     string `gorm:"type:varchar(7)" json:"color_code"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
