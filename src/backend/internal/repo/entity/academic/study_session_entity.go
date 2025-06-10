package academic

import "time"

type Study_session struct {
	Session_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"session_id"`

	Student_id string    `gorm:"type:uuid; index" json:"student_id"`
	Student    []Student `gorm:"foreignKey:Student_id"`

	Subject_id string  `gorm:"type:uuid; index" josn:"subject_id"`
	Subject    Subject `gorm:"foreignKey:Subject_id"`

	Title        string    `gorm:"type:varchar(100); not null" json:"title"`
	Duration     int       `gorm:"type:int; not null" json:"duration"`
	Notes        string    `gorm:"type:text; not null" json:"notes"`
	Session_date time.Time `gorm:"type:timestamp; not null" json:"session_date"`
}
