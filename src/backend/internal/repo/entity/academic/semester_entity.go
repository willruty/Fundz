package academic

import "time"

type Semester struct {
	Semester_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"semester_id"`

	Student_id string  `gorm:"type:uuid; index" json:"student_id"`
	Student    []Student `gorm:"foreignKey:Student_id"`

	Start_date time.Time `gorm:"type:date; not null" json:"start_date"`
	End_date   time.Time `gorm:"type:date; not null" json:"end_date"`
}
