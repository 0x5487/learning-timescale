package main

import (
	"candlestick/cmd"

	_ "github.com/lib/pq"
)

func main() {
	cmd.Execute()
}
