package cmd

import (
	"context"
	"fmt"

	"os"
	"time"

	"candlestick/cmd/testdata"

	"github.com/gofiber/fiber/v2"
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
	},
}

func startup(ctx context.Context) error {
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
