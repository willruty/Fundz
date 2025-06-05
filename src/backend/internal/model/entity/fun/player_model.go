package entity

import "time"

type Player struct {
	Nickname   string    `gorm:"type:varchar(100);not null" json:"nickname" binding:"required"`

	// user_id

	Updated_at time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
	Created_at time.Time `gorm:"type:timestamp(0)" json:"created_at"`
}
