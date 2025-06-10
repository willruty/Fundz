package finance

import "github.com/shopspring/decimal"

type MostTransactioMadeDTO struct {
	// transaction id
	Count       int64
	Receiver    string
	Avg_spent   float64
	Total_spent decimal.Decimal
}
