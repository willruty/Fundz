package academic

type AssisgnmentStatsDTO struct {
	// subject id
	Subject_name string          `json:"subject_name"`
	Delivered int64 `json:"delivered"`
	Pending int64 `json:"pending"`
	Delayed int64 `json:"delayed"`
}