package cmd

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"os"
	"time"

	"learning-timescaledb/cmd/testdata"
	"learning-timescaledb/internal/pkg/initialize"
	"learning-timescaledb/pkg/market/delivery/http"
	"learning-timescaledb/pkg/market/repository/timescaledb"
	"learning-timescaledb/pkg/market/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/nite-coder/blackbear/pkg/config"
	"github.com/nite-coder/blackbear/pkg/log"
	"github.com/spf13/cobra"
)

var (
	app      *fiber.App
	ConfPath string
)

func init() {
	rootCmd.Flags().StringVar(&ConfPath, "config", "app.yaml", "config file (default is ./app.yaml)")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "candlestick",
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

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		<-quit

		err = shutdown(ctx)
		if err != nil {
			log.Err(err).Warnf("cmd: fail to shutdown")
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

	// fiber web engine
	app, err = initialize.Fiber()
	if err != nil {
		return err
	}

	// repo
	tradeRepo := timescaledb.NewTradeRepo(db)

	// usecase
	candlestickUC := usecase.NewCandlestickUsecase(tradeRepo)

	// handlers
	candlestickHandler := http.NewCandlestickHandler(candlestickUC)
	err = http.RegisterRoute(ctx, app, candlestickHandler)
	if err != nil {
		return err
	}

	bind, _ := config.String("fiber.bind", ":80")
	go func() {
		err := app.Listen(bind)
		if err != nil {
			log.Err(err).Fatal("fail to start fiber engine")
		}
	}()

	log.Infof("candlestick started. bind at: %s", bind)
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
			err := app.Shutdown()
			if err != nil {
				return err
			}

			log.Info("service was shutdown successfully")
			return nil
		}
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(testdata.CMD)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
