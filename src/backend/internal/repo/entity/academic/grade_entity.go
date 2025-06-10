package academic

import "time"

type Grade struct {
	Grade_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"grade_id"`
	// student id
	// subject id
	Title      string    `gorm:"type:varchar(100); not null" json:"title"`
	Value      float64   `gorm:"type:float; not null" json:"value"` // Assuming grade is a float value
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
