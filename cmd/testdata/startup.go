package testdata

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"learning-timescaledb/internal/pkg/initialize"
	"learning-timescaledb/pkg/domain"
	"learning-timescaledb/pkg/market/repository/timescaledb"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

// CMD represents the base command when called without any subcommands
var CMD = &cobra.Command{
	Use: "testdata",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		defer log.Flush()
		defer func() {
			if r := recover(); r != nil {
				// unknown error
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("unknown error: %v", r)
				}
				log.Err(err).Panic("unknown error")
			}

		}()

		err := startup(ctx)
		if err != nil {
			log.Panicf("cmd: fail to startup: %v", err)
			return
		}

	},
}

func startup(ctx context.Context) error {
	// config
	err := initialize.Config("")
	if err != nil {
		return err
	}

	// log
	err = initialize.Logger()
	if err != nil {
		return err
	}

	// timescaledb
	db, err := initialize.TimescaleDB(ctx)
	if err != nil {
		return err
	}

	// load test data
	trades, err := load()
	if err != nil {
		return err
	}

	tradeRepo := timescaledb.NewTradeRepo(db)

	now := time.Now()
	err = tradeRepo.BulkInsert(ctx, trades, 10000)
	if err != nil {
		return err
	}

	duration := time.Since(now)
	log.Str("duration", duration.String()).Info("done!!")
	return nil
}

func shutdown(ctx context.Context) error {
	log.Info("candlestick is shutting down...")

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:

			log.Info("service was shutdown successfully")
			return nil
		}
	}
}

func load() ([]*domain.Trade, error) {
	log.Info("loading test data...")

	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "test", "btc_usdt.csv")

	file, err := os.OpenFile(path, os.O_RDONLY, 0777) // os.O_RDONLY 表示只讀、0777 表示(owner/group/other)權限
	if err != nil {
		return nil, err
	}

	//result := make([]domain.Trade, 4000000)
	result := []*domain.Trade{}

	// read
	r := csv.NewReader(file)
	r.Comma = ',' // 以何種字元作分隔，預設為`,`。所以這裡可拿掉這行
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		f, err := strconv.ParseFloat(strings.TrimSpace(record[0]), 64)
		if err != nil {
			log.Warnf("can't parse time column: %s", record[0])
			continue
		}

		t := f * 1000000

		side, err := strconv.Atoi(strings.TrimSpace(record[4]))
		if err != nil {
			log.Warnf("can't parse side column: %s", record[4])
			continue
		}

		trade := &domain.Trade{
			Time:   time.UnixMicro(int64(t)),
			ID:     strings.TrimSpace(record[1]),
			Market: "BTC_USDT",
			Price:  decimal.RequireFromString(strings.TrimSpace(record[2])),
			Size:   decimal.RequireFromString(strings.TrimSpace(record[3])),
			Side:   domain.Side(side),
		}

		trade.Volume = trade.Price.Mul(trade.Size)

		result = append(result, trade)
	}

	log.Infof("loaded: %d", len(result))

	return result, nil
}
