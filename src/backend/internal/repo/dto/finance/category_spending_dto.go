package finance

import "github.com/shopspring/decimal"

type CategorySpendingDTO struct {
	// category id
	Category_name string
	Total_spent   decimal.Decimal
}
