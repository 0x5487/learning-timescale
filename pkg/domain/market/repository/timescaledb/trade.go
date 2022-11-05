package timescaledb

import (
	"candlestick/pkg/domain"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TradeRepo struct {
	db *pgxpool.Pool
}

func NewTradeRepo(db *pgxpool.Pool) *TradeRepo {
	return &TradeRepo{
		db: db,
	}
}

func (repo *TradeRepo) BulkInsert(ctx context.Context, trades []*domain.Trade) error {

	rows := [][]interface{}{}

	for _, trade := range trades {
		rows = append(rows, []interface{}{trade.Time, trade.OrderID, trade.Market, trade.Side, trade.Price, trade.Size})
	}

	_, err := repo.db.CopyFrom(ctx,
		pgx.Identifier{"trade"},
		[]string{"time", "order_id", "market", "side", "price", "size"},
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return err
	}

	return nil
}
