package fun

type MatchPlayer struct {
	Player_id string `gorm:"type:uuid; primaryKey" json:"player_id"`
	Player    Player `gorm:"foreignKey:Player_id"`

	Match_id string `gorm:"type:uuid; primaryKey" json:"match_id"`
	Match    Match  `gorm:"foreignKey:Match_id"`

	Team      string `gorm:"type:varchar(50)" json:"team"`
	Score     int64  `gorm:"type:bigint; not null; default:0;" json:"score"`
	Is_winner bool   `gorm:"type:boolean; not null; default:false;" json:"is_winner"`
}