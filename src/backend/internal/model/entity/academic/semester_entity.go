package academic

import "time"

type Semester struct {
	SemesterId string    `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"semester_id"`
	StudentId  string    `gorm:"type:uuid; index" json:"student_id"`
	StartDate  time.Time `gorm:"type:date; not null" json:"start_date"`
	EndDate    time.Time `gorm:"type:date; not null" json:"end_date"`
}
