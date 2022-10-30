package main

import (
	"candlestick/cmd"

	_ "github.com/jackc/pgx/v5"
)

func main() {
	cmd.Execute()
}
