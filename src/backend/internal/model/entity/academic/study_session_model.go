package entity

import (
	"time"
)

type Study_session struct {
	Ss_id string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"ss_id" binding:"required"`

	// user_id
	// subject_id

	Ss_duration_minutes int       `gorm:"type:integer;not null" json:"ss_duration_minutes" binding:"required"`
	Ss_date             time.Time `gorm:"type:timestamp;not null" json:"ss_date" binding:"required"`
	Ss_description      string    `gorm:"type:varchar(300)" json:"ss_description"`

	Created_at time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
}
