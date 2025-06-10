package finance

import (
	"time"

	"github.com/shopspring/decimal"
)

type DaySummary struct {
	Date         time.Time
	Total_income decimal.Decimal
	Total_expense decimal.Decimal
}
