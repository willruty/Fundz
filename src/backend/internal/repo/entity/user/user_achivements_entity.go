package usuario

type UserAchievement struct {
	Achivement_id       string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"achivement_id"`
	Achivement_name     string `gorm:"type:varchar(100);not null" json:"achivement_name"`
	Achivement_category string `gorm:"type:varchar(50);not null" json:"achivement_category"`
	Points              int64  `gorm:"type:bigint; not null; default:0;" json:"points"`
	Icon_url            string `gorm:"type:text; not null" json:"icon_url"`
}
