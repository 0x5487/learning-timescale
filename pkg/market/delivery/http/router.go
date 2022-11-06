package http

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(ctx context.Context, app *fiber.App, handler *CandlestickHandler) error {
	app.Get("/spot/candlestick", handler.listCandlestickEndpoint)

	return nil
}
