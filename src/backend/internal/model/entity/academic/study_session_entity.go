package academic

import "time"

type Study_session struct {
	SessionId   string    `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"session_id"`
	StudentId   string    `gorm:"type:uuid; index" json:"student_id"`
	SubjectId   string    `gorm:"type:uuid; index" josn:"subject_id"`
	Title       string    `gorm:"type:varchar(100); not null" json:"title"`
	Duration    int       `gorm:"type:int; not null" json:"duration"`
	Notes       string    `gorm:"type:text; not null" json:"notes"`
	SessionDate time.Time `gorm:"type:timestamp; not null" json:"session_date"`
}
