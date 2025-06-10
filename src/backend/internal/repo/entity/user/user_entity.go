package user

type User struct {
	User_id       string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"user_id"`
	User_name     string `gorm:"type:varchar(70);not null" json:"user_name"`
	User_email    string `gorm:"type:varchar(70);not null" json:"user_email"`
	User_password string `gorm:"type:varchar(30); not null" json:"user_password"`
}
