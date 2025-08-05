package usuario

type UserAccount struct {
	UserId   string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"user_id"`
	Name     string `gorm:"type:varchar(70);not null" json:"name"`
	Email    string `gorm:"type:varchar(70);not null" json:"email"`
	Password string `gorm:"type:varchar(30); not null" json:"password"`
}

func (*UserAccount) TableName() string {
	return "user"
}
