package finance

import "github.com/shopspring/decimal"

type TopEarnerGoalDTO struct {
	// goal id
	Goal_name       string
	Received_amount decimal.Decimal
	Percentage_of_target float64
}