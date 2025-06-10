package academic

import "time"

type Grade struct {
	Grade_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"grade_id"`

	Student_id string    `gorm:"type:uuid; index" json:"student_id"`
	Student    []Student `gorm:"foreignKey:Student_id"`

	Subject_id string  `gorm:"type:uuid; index" josn:"subject_id"`
	Subject    Subject `gorm:"foreignKey:Subject_id"`

	Title      string    `gorm:"type:varchar(100); not null" json:"title"`
	Value      float64   `gorm:"type:float; not null" json:"value"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
