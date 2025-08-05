package fun

import (
	"time"
)

type Player struct {
	PlayerId   string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"player_id"`
	UserId     string `gorm:"type:uuid; primaryKey" json:"user_id"`
	Nickname   string `gorm:"type:varchar(50);not null" json:"nickname"`
	Rank       int64  `gorm:"type:bigint; not null; default:0;" json:"rank"`
	Level      int64  `gorm:"type:bigint; not null; default:0;" json:"level"`
	Avatar_img string `gorm:"type:text; not null" json:"avatar_img"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
