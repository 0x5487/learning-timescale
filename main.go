package main

import (
	"learning-timescaledb/cmd"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	cmd.Execute()
}
