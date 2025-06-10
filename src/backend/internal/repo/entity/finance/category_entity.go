package finance

import "time"

type Category struct {
	Category_id string `gorm:"type:varchar(50);primaryKey" json:"category_id"` // unique identifier for the category
	// account id
	Category_title string `gorm:"type:varchar(50);not null" json:"category_title"` // title of the category
	Description    string `gorm:"type:varchar(255)" json:"description"`            // description of the category
	Color_code     string `gorm:"type:varchar(7)" json:"color_code"`               // color code for the category, e.g., "#FF5733"

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"` // timestamp when the category was created
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"` // timestamp when the category was last updated
}
