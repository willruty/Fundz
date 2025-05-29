package entity

import (
	"time"

	"github.com/google/uuid"
)

type Score struct {
	Score_id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"score_id" binding:"required"`

	// game_id
	// player_id

	Score    int `gorm:"type:integer;not null" json:"score" binding:"required"`
	Position int `gorm:"type:integer;not null" json:"position" binding:"required"`

	Created_at time.Time `gorm:"type:timestamp(0)" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp(0)" json:"updated_at"`
}
