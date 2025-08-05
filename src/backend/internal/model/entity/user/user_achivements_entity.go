package usuario

type UserAchievement struct {
	AchivementId       string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"achivement_id"`
	AchivementName     string `gorm:"type:varchar(100);not null" json:"achivement_name"`
	AchivementCategory string `gorm:"type:varchar(50);not null" json:"achivement_category"`
	Points             int64  `gorm:"type:bigint; not null; default:0;" json:"points"`
	IconUrl            string `gorm:"type:text; not null" json:"icon_url"`
}
