package timescaledb

import (
	"context"
	"learning-timescaledb/pkg/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nite-coder/blackbear/pkg/log"
)

type TradeRepo struct {
	db *pgxpool.Pool
}

func NewTradeRepo(db *pgxpool.Pool) *TradeRepo {
	return &TradeRepo{
		db: db,
	}
}

func (repo *TradeRepo) BulkInsert(ctx context.Context, trades []*domain.Trade, batchSize int32) error {
	logger := log.FromContext(ctx)

	rows := [][]interface{}{}

	for idx, trade := range trades {
		idx++

		rows = append(rows, []interface{}{trade.Time, trade.ID, trade.Market, trade.Side, trade.Price, trade.Size, trade.Volume})

		if idx%int(batchSize) == 0 || idx == len(trades) {
			_, err := repo.db.CopyFrom(ctx,
				pgx.Identifier{"trade"},
				[]string{"time", "id", "market", "side", "price", "size", "volume"},
				pgx.CopyFromRows(rows),
			)

			rows = [][]interface{}{}

			if err != nil {
				return err
			}

			logger.Debugf("write count: %d", idx)
		}
	}

	return nil
}
