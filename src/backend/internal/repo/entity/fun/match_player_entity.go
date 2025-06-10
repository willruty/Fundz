package fun

type MatchPlayer struct {
	// player id
	// match id
	Team string `gorm:"type:varchar(50)" json:"team"`
	Score int64  `gorm:"type:bigint; not null; default:0;" json:"score"`
	Is_winner bool   `gorm:"type:boolean; not null; default:false;" json:"is_winner"`
}
