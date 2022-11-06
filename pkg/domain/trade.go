package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Side int8

const (
	Side_Default Side = 0
	Side_Sell    Side = 1
	Side_Buy     Side = 2
)

type Trade struct {
	Time    time.Time
	ID string
	Market  string
	Side    Side
	Price   decimal.Decimal
	Size    decimal.Decimal
	Volume  decimal.Decimal
}
