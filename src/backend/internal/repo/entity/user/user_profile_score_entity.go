package user

type UserProfileScore struct {
	Social_level            int64   `gorm:"type:bigint; not null; default:0;" json:"social_level"`
	Financial_control_level int64   `gorm:"type:bigint; not null; default:0;" json:"financial_control_level"`
	Nerd_level              int64   `gorm:"type:bigint; not null; default:0;" json:"nerd_level"`
	Party_level             int64   `gorm:"type:bigint; not null; default:0;" json:"party_level"`
	Overall_level           float64 `gorm:"type:double precision; not null; default:0;" json:"overall_level"`
	XP                      float64 `gorm:"type:double precision; not null; default:0;" json:"xp"`

	User_id     string      `gorm:"type:uuid; primaryKey" json:"user_id"`
	UserAccount UserAccount `gorm:"foreignKey:User_id;references:User_id" json:"user_account"`
	Badges      []UserBadge `gorm:"many2many:user_profile_score_badges;" json:"badges"`
}
