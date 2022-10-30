package initialize

import (
	"database/sql"

	"github.com/nite-coder/blackbear/pkg/config"
)

func TimescaleDB() (*sql.DB, error) {
	connStr, err := config.String("timescaledb.connection_string")
	if err != nil {
		return nil, err
	}

	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
