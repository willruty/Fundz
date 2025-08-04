package finance

import "time"

type Goal struct {
	Goal_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"goal_id"`

	Account_id string  `gorm:"type:uuid; primaryKey" json:"account_id"`
	Account    Account `gorm:"foreignKey:Account_id; references:Account_id" json:"user_account"`

	Category_id string   `gorm:"type:uuid; primaryKey" json:"category_id"`
	Category    Category `gorm:"foreignKey:Category_id; references:Category_id" json:"category"`

	Target_amount  float64   `gorm:"type:decimal(10,2);not null" json:"target_amount"`
	Current_amount float64   `gorm:"type:decimal(10,2);not null" json:"current_amount"`
	Goal_title     string    `gorm:"type:varchar(100);not null" json:"goal_title"`
	Description    string    `gorm:"type:varchar(255)" json:"description"`
	Deadline       time.Time `gorm:"type:timestamp" json:"deadline"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
