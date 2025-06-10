package user

import "time"

type UserBadge struct {
	Badge_id         string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"badge_id"`
	Badge_name       string    `gorm:"type:varchar(100); not null" json:"badge_name"`
	Category         string    `gorm:"type:varchar(50); not null" json:"category"`
	Unlock_condition string    `gorm:"type:text; not null" json:"unlock_condition"`
	Icon_url         string    `gorm:"type:text; not null" json:"icon_url"`
	Created_at       time.Time `gorm:"type:timestamp; not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
