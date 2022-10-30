package timescaledb

import (
	"candlestick/pkg/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TradeRepo struct {
	db *pgxpool.Pool
}

func NewTradeRepo() *TradeRepo {
	return &TradeRepo{}
}

func (repo *TradeRepo) Bulk(ctx context.Context, trades []*domain.Trade) error {

	// rows := [][]interface{}{
	// 	{"John", "Smith", int32(36)},
	// 	{"Jane", "Doe", int32(29)},
	// }

	// copyCount, err := repo.db.CopyFrom(
	// 	pgx.Identifier{"trade"},
	// 	[]string{"first_name", "last_name", "age"},
	// 	pgx.CopyFromRows(rows),
	// )

	return nil
}
