package finance

import "github.com/shopspring/decimal"

type MonthSummaryDTO struct {
	// account id
	Month                   int
	Total_income            decimal.Decimal
	Total_expense           decimal.Decimal
	Net_balance             decimal.Decimal
	Goal_completion_percent float64
}
