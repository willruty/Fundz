package academic

import "time"

type Assignment struct {
	AssignmentId string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"assignment_id"`
	StudentId    string `gorm:"type:uuid; index" json:"student_id"`
	SubjectId    string `gorm:"type:uuid; index" josn:"subject_id"`
	Title        string `gorm:"type:varchar(100); not null" json:"title"`
	Description  string `gorm:"type:text; not null" json:"description"`
	Deadline     string `gorm:"type:timestamp; not null" json:"deadline"`
	Status       string `gorm:"type:varchar(20); not null; default:'pending'" json:"status"`
	IsDelivered  bool   `gorm:"type:boolean; not null; default:false" json:"is_delivered"`

	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
