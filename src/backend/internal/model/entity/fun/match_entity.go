package fun

import "time"

type Match struct {
	MatchId    string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"match_id"`
	GameTypeId string `gorm:"type:uuid; primaryKey" json:"game_type_id"`
	MatchName  string `gorm:"type:varchar(50);not null" json:"match_name"`
	Location   string `gorm:"type:varchar(100);not null" json:"location"`
	PlayedAt   string `gorm:"type:timestamp; default:current_timestamp" json:"played_at"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
