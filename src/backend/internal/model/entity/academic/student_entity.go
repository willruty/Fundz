package academic

import (
	usuario "fundz/internal/model/entity/user"
	"time"
)

type Student struct {
	Student_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"student_id"`

	User_id     string              `gorm:"type:uuid; index" json:"user_id"`
	UserAccount usuario.UserAccount `gorm:"foreignKey:User_id;references:User_id" json:"user_account"`

	Course    string `gorm:"type:varchar(70); not null" json:"course"`
	Institute string `gorm:"type:varchar(100)" json:"institute"`
	Semester  int64  `gorm:"type:bigint; not null; default:1;" json:"semester"`

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
