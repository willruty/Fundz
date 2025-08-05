package academic

import (
	"github.com/shopspring/decimal"
	"time"
)

type Grade struct {
	GradeId   string          `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"grade_id"`
	StudentId string          `gorm:"type:uuid; index" json:"student_id"`
	SubjectId string          `gorm:"type:uuid; index" josn:"subject_id"`
	Title     string          `gorm:"type:varchar(100); not null" json:"title"`
	Value     decimal.Decimal `gorm:"type:float; not null" json:"value"`
	UpdatedAt time.Time       `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
