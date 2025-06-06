package entity

type Game_type struct {
	Game_type_id          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;not null" json:"game_type_id" binding:"required"`
	Game_type_name        string `gorm:"type:varchar(300);not null" json:"game_type_name" binding:"required"`
	Game_type_description string `gorm:"type:varchar(300)" json:"game_type_description"`
}
