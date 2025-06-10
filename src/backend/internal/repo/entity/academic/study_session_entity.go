package academic

import "time"

type Study_session struct {
	Session_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"session_id"`
	// student id
	// subject id
	Title        string    `gorm:"type:varchar(100); not null" json:"title"`
	Duration     int       `gorm:"type:int; not null" json:"duration"` // Duration in minutes
	Notes        string    `gorm:"type:text; not null" json:"notes"`
	Session_date time.Time `gorm:"type:timestamp; not null" json:"session_date"` // Date of the study session
}
