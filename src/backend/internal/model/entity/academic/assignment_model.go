package entity

import (
	"time"

	"github.com/shopspring/decimal"
)

type Assignment struct {
	Assignment_id          string          `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"assignment_id" binding:"required"`
	Assignment_title       string          `gorm:"type:varchar(250);not null" json:"assignment_title" binding:"required"`
	Assignment_description string          `gorm:"type:varchar(250)" json:"assignment_description"`
	Assignment_grade       decimal.Decimal `gorm:"type:numeric(250,2)" json:"assignment_grade"`
	Assignment_status      string          `gorm:"type:varchar(50);not null" json:"assignment_status" binding:"required"`

	// subject_id

	Updated_at time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
	Created_at time.Time `gorm:"type:timestamp(0)" json:"created_at"`

	Delivery_date time.Time `gorm:"type:timestamp;not null" json:"delivery_date" binding:"required"`
}
