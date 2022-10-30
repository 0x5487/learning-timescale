package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trade struct {
	Time    time.Time
	OrderID string
	Market  string
	Side    int8
	Price   decimal.Decimal
	Size    decimal.Decimal
}
