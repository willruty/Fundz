package academic

import "time"

type Subject struct {
	SubjectId   string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"subject_id"`
	StudentId   string `gorm:"type:uuid; index" json:"student_id"`
	SemesterId  string `gorm:"type:uuid; index" json:"semester_id"`
	ProfessorId string `gorm:"type:uuid; index" json:"professor_id"`
	Name        string `gorm:"type:varchar(50);not null" json:"name"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
