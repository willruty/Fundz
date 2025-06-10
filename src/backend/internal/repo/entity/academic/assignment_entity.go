package academic

import "time"

type Assignment struct {
	Assignment_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"assignment_id"`
	// student id
	// subject id
	Title        string `gorm:"type:varchar(100); not null" json:"title"`
	Description  string `gorm:"type:text; not null" json:"description"`
	Deadline     string `gorm:"type:timestamp; not null" json:"deadline"`
	Status       string `gorm:"type:varchar(20); not null; default:'pending'" json:"status"`
	Is_delivered bool   `gorm:"type:boolean; not null; default:false" json:"is_delivered"`

	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
