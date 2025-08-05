package fun

type MatchPlayer struct {
	PlayerId string `gorm:"type:uuid; primaryKey" json:"player_id"`
	MatchId  string `gorm:"type:uuid; primaryKey" json:"match_id"`
	Team     string `gorm:"type:varchar(50)" json:"team"`
	Score    int64  `gorm:"type:bigint; not null; default:0;" json:"score"`
	IsWinner bool   `gorm:"type:boolean; not null; default:false;" json:"is_winner"`
}
