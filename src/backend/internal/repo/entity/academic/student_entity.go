package academic

import "time"

type Student struct {
	Student_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"student_id"`
	// user id
	Course    string `gorm:"type:varchar(70); not null" json:"course"`
	Institute string `gorm:"type:varchar(100)" json:"institute"`
	Semester  int64  `gorm:"type:bigint; not null; default:1;" json:"semester"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time   `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
