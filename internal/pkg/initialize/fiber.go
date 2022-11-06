package initialize

import (
	"errors"
	"learning-timescaledb/pkg/domain"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nite-coder/blackbear/pkg/config"
	"github.com/nite-coder/blackbear/pkg/log"
)

func Fiber() (*fiber.App, error) {
	prefork, _ := config.Bool("fiber.prefork", false)
	readTimeout, _ := config.Duration("fiber.read_timeout", time.Duration(60)*time.Second)
	writeTimeout, _ := config.Duration("fiber.read_timeout", time.Duration(60)*time.Second)

	app := fiber.New(fiber.Config{
		Prefork:      prefork,
		ErrorHandler: errorHandler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	})

	app.Use(recover.New(recover.Config{EnableStackTrace: true, StackTraceHandler: stackTraceHandler}))
	app.Use(requestid.New())
	app.Use(setRequestID)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})

	return app, nil
}

var errorHandler = func(c *fiber.Ctx, err error) error {
	ctx := c.UserContext()
	logger := log.FromContext(ctx)

	var appErr *domain.AppError
	if errors.As(err, &appErr) {
		logger.Err(appErr).Debug("fiber: error handled")
		return c.Status(appErr.HTTPStatus).JSON(appErr)
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		if fiberErr.Code == http.StatusNotFound {
			return c.Status(http.StatusNotFound).Send(nil)
		}
	}

	logger.Err(err).Error("fiber: unknown error")
	return c.Status(fiber.StatusInternalServerError).JSON(domain.ErrInternal)
}

var setRequestID = func(c *fiber.Ctx) error {
	ctx := c.UserContext()
	logger := log.FromContext(ctx)

	rid := c.Locals("requestid")
	ctx = logger.Any("request_id", rid).WithContext(ctx)

	c.SetUserContext(ctx)
	return c.Next()
}

var stackTraceHandler = func(c *fiber.Ctx, e interface{}) {
	ctx := c.UserContext()
	logger := log.FromContext(ctx)
	logger.Errorf("panic: %v", e)
}
