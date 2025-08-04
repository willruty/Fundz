package finance

import (
	"time"

	"github.com/shopspring/decimal"
)

type GoalProgress struct {
	// goal id
	Goal_name        string
	Target_amount    decimal.Decimal
	Current_amount   decimal.Decimal
	Progress_percent float64
	Deadline         time.Time
}
