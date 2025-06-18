package fun

import "time"

type Match struct {
	Match_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"match_id"`

	Game_type_id string    `gorm:"type:uuid; primaryKey" json:"game_type_id"`
	Game_type    GameType `gorm:"foreignKey:Game_type_id"`

	Match_name string `gorm:"type:varchar(50);not null" json:"match_name"`
	Location   string `gorm:"type:varchar(100);not null" json:"location"`
	Played_at  string `gorm:"type:timestamp; default:current_timestamp" json:"played_at"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
