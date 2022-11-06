package http

import (
	"learning-timescaledb/pkg/domain"

	"github.com/gofiber/fiber/v2"
)

type CandlestickHandler struct {
	usecase domain.CandlestickUsecase
}

func NewCandlestickHandler(usecase domain.CandlestickUsecase) *CandlestickHandler {
	return &CandlestickHandler{
		usecase: usecase,
	}
}

func (h *CandlestickHandler) listCandlestickEndpoint(c *fiber.Ctx) error {

	return c.Status(200).SendString("hello")
}
