package academic

import "time"

type Assignment struct {
	Assignment_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"assignment_id"`

	Student_id string    `gorm:"type:uuid; index" json:"student_id"`
	Student    []Student `gorm:"foreignKey:Student_id"`

	Subject_id string  `gorm:"type:uuid; index" josn:"subject_id"`
	Subject    Subject `gorm:"foreignKey:Subject_id"`

	Title        string `gorm:"type:varchar(100); not null" json:"title"`
	Description  string `gorm:"type:text; not null" json:"description"`
	Deadline     string `gorm:"type:timestamp; not null" json:"deadline"`
	Status       string `gorm:"type:varchar(20); not null; default:'pending'" json:"status"`
	Is_delivered bool   `gorm:"type:boolean; not null; default:false" json:"is_delivered"`

	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
