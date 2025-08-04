package academic

type SemesterStatsDTO struct {
	// semester id 
	Semester_name string `json:"semester_name"`
	Avg_grade    float64 `json:"avg_grade"`
	Total_subjects int     `json:"total_subjects"`
	Passed_subjects int    `json:"passed_subjects"`
	Failed_subjects int    `json:"failed_subjects"`
}