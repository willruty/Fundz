package entity

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	Game_id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"game_id" binding:"required"`

	// game_type_id

	Game_date     time.Time `gorm:"type:timestamp;not null" json:"game_date" binding:"required"`
	Game_location string    `gorm:"type:varchar(300)" json:"game_location"`
	
	Created_at    time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	Updated_at    time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
}
