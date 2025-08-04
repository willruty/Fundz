package academic

type GradeStatsDTO struct {
	Grade_id     string  `json:"grade_id"`
	Title		string  `json:"title"`
	Type		 string  `json:"type"`
	Score		 float64 `json:"score"`
	Max_score	 float64 `json:"max_score"`
	Weight		 float64 `json:"weight"`
}
