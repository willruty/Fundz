package fun

import "time"

type GameType struct {
	Game_type_id   string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"game_type_id"`
	Game_type_name string `gorm:"type:varchar(50);not null" json:"game_type_name"`
	Description    string `gorm:"type:text; not null" json:"description"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time   `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
