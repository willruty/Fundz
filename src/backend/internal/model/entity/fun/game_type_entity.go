package fun

import "time"

type GameType struct {
	GameTypeId   string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"game_type_id"`
	GameTypeName string `gorm:"type:varchar(50);not null" json:"game_type_name"`
	Description  string `gorm:"type:text; not null" json:"description"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
