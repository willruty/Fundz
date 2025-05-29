package entity

import (
	"time"

	"github.com/google/uuid"
)

type Subject struct {
	Subject_id   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"subject_id" binding:"required"`
	Subject_name string    `gorm:"type:varchar(250);not null" json:"subject_name" binding:"required"`

	// user_id

	Professor_name string `gorm:"type:varchar(250);not null" json:"professor_name" binding:"required"`
	Missed_days    int    `gorm:"type:integer" json:"missed_days"`
	Semester       int    `gorm:"type:integer;not null" json:"semester" binding:"required"`

	Created_at time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
}
