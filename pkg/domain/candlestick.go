package domain

import "context"

type FindCandlestickOptions struct {
	Market   string
	Limit    int32
	From     int64
	To       int64
	Interval string
}

type CandlestickUsecase interface {
	Candlesticks(ctx context.Context, opts *FindCandlestickOptions) ([]*Candlestick, error)
}

type TradeRepository interface {
	Candlesticks(ctx context.Context, opts *FindCandlestickOptions) ([]*Candlestick, error)
	BulkInsert(ctx context.Context, trades []*Trade, batchSize int32) error
}

type KLine1MinuteRepository interface {
	Candlesticks(ctx context.Context, opts *FindCandlestickOptions) ([]*Candlestick, error)
}

type KLine1HourRepository interface {
	Candlesticks(ctx context.Context, opts *FindCandlestickOptions) ([]*Candlestick, error)
}

type KLine1DayRepository interface {
	Candlesticks(ctx context.Context, opts *FindCandlestickOptions) ([]*Candlestick, error)
}

type Candlestick struct {
	Time        int64
	Open        string
	High        string
	Low         string
	Close       string
	VolumeBase  string
	VolumeQuote string
}
