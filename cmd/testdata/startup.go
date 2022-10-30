package testdata

import (
	"candlestick/internal/pkg/initialize"
	"context"
	"fmt"
	"time"

	"github.com/nite-coder/blackbear/pkg/log"
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
	_, err = initialize.TimescaleDB()
	if err != nil {
		return err
	}

	log.Info("test data !!!")
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
