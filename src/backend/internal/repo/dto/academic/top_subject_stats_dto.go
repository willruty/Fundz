package academic

type TopSubjectStatsDTO struct {
	// subject id
    Subject_name string  `gorm:"column:subject_name"` 
    Avg_grade    float64 `gorm:"column:avg_grade"`
    Study_time   int64   `gorm:"column:study_time"`
    Total_grades int64   `gorm:"column:total_grades"`
}
