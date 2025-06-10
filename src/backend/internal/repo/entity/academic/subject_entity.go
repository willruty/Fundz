package academic

import "time"

type Subject struct {
	Subject_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"subject_id"`

	Student_id string  `gorm:"type:uuid; index" json:"student_id"`
	Student    []Student `gorm:"foreignKey:Student_id"`

	Semester_id string   `gorm:"type:uuid; index" json:"semester_id"`
	Semester    []Semester `gorm:"foreignKey:Semester_id"`

	Professor_id string    `gorm:"type:uuid; index" json:"professor_id"`
	Professor    Professor `gorm:"foreignKey:Professor_id"`

	Name string `gorm:"type:varchar(50);not null" json:"name"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
