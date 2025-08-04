package academic

type SubjectStatsDTO struct {
	// subject id
	Subject_name string          `json:"subject_name"`
	Grades       []GradeStatsDTO `json:"grades"`
}
