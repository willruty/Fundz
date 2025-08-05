package academic

import (
	"time"
)

type Student struct {
	StudentId string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"student_id"`
	UserId    string `gorm:"type:uuid; index" json:"user_id"`
	Course    string `gorm:"type:varchar(70); not null" json:"course"`
	Institute string `gorm:"type:varchar(100)" json:"institute"`
	Semester  int64  `gorm:"type:bigint; not null; default:1;" json:"semester"`

	CreatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
