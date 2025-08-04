package academic

type Professor struct {
	Professor_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"professor_id"`
	Name         string `gorm:"type:varchar(100); not null" json:"name"`
	Email        string `gorm:"type:varchar(50); unique" json:"email"`
	Phone        string `gorm:"type:varchar(20)" json:"phone"`
}
