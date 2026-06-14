package middleware

import (
	"trainers-manager/pkg/logger"

	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(l logger.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		traceID, _ := c.Locals(TraceIDKey).(string)

		l.Info("request trace_id=%s method=%s path=%s status=%d duration=%s",
			traceID,
			c.Method(),
			c.Path(),
			c.Response().StatusCode(),
			time.Since(start),
		)
		return err
	}
}
