package finance

import "time"

type Goal struct {
	Goal_id string `gorm:"type:uuid; default:uuid_generate_v4(); primaryKey" json:"goal_id"`
	// account id
	Target_amount  float64 `gorm:"type:decimal(10,2);not null" json:"target_amount"`  // target amount to reach
	Current_amount float64 `gorm:"type:decimal(10,2);not null" json:"current_amount"` // current amount saved towards the goal
	Goal_title     string  `gorm:"type:varchar(100);not null" json:"goal_title"`      // title of the goal
	// category id
	Description string    `gorm:"type:varchar(255)" json:"description"` // description of the goal
	Due_date    time.Time `gorm:"type:timestamp" json:"due_date"`       // due date for the goal

	Created_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"created_at"`
	Updated_at time.Time `gorm:"type:timestamp; default:current_timestamp" json:"updated_at"`
}
