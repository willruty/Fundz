package usuario

type UserProfileScore struct {
	SocialLevel           int64   `gorm:"type:bigint; not null; default:0;" json:"social_level"`
	FinancialControlLevel int64   `gorm:"type:bigint; not null; default:0;" json:"financial_control_level"`
	NerdLevel             int64   `gorm:"type:bigint; not null; default:0;" json:"nerd_level"`
	PartyLevel            int64   `gorm:"type:bigint; not null; default:0;" json:"party_level"`
	OverallLevel          float64 `gorm:"type:double precision; not null; default:0;" json:"overall_level"`
	XP                    float64 `gorm:"type:double precision; not null; default:0;" json:"xp"`

	UserId string `gorm:"type:uuid; primaryKey" json:"user_id"`
}
