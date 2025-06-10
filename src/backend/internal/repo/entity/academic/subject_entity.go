package academic

import "time"

type Subject struct {
	Subject_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"subject_id"`
	// student id
	// semester id
	Name string `gorm:"type:varchar(50);not null" json:"name"`
	// professor id
	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
