package usuario

import "time"

type UserBadge struct {
	BadgeId         string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"badge_id"`
	BadgeName       string    `gorm:"type:varchar(100); not null" json:"badge_name"`
	Category        string    `gorm:"type:varchar(50); not null" json:"category"`
	UnlockCondition string    `gorm:"type:text; not null" json:"unlock_condition"`
	IconUrl         string    `gorm:"type:text; not null" json:"icon_url"`
	CreatedAt       time.Time `gorm:"type:timestamp; not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}
