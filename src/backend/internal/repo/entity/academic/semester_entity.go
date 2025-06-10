package academic

import "time"

type Semester struct {
	// student id
	Start_date time.Time `gorm:"type:date; not null" json:"start_date"`
	End_date   time.Time `gorm:"type:date; not null" json:"end_date"`
}
