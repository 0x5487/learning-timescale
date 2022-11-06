package usecase

import (
	"context"
	"learning-timescaledb/pkg/domain"
)

type CandlestickUsecase struct {
	db domain.TradeRepository
}

func NewCandlestickUsecase(db domain.TradeRepository) *CandlestickUsecase {
	return &CandlestickUsecase{
		db: db,
	}
}

func (uc *CandlestickUsecase) Candlesticks(ctx context.Context, opts *domain.FindCandlestickOptions) ([]*domain.Candlestick, error) {
	return nil, nil
}
