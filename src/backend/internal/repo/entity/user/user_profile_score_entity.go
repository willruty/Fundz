package user

type UserProfileScore struct {
	
	// user_id
	Social_level            int64   `gorm:"type:bigint; not null; default:0;" json:"social_level"`
	Financial_control_level int64   `gorm:"type:bigint; not null; default:0;" json:"financial_control_level"`
	Nerd_level              int64   `gorm:"type:bigint; not null; default:0;" json:"nerd_level"`
	Party_level             int64   `gorm:"type:bigint; not null; default:0;" json:"party_level"`
	Overall_level           float64 `gorm:"type:double precision; not null; default:0;" json:"overall_level"`
	XP                      float64 `gorm:"type:double precision; not null; default:0;" json:"xp"`

	// badges []badges
}
