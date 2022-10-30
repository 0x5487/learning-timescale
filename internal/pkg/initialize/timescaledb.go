package initialize

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nite-coder/blackbear/pkg/config"
)

func TimescaleDB(ctx context.Context) (*pgxpool.Pool, error) {
	connStr, err := config.String("timescaledb.connection_string")
	if err != nil {
		return nil, err
	}

	// Connect to database
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
