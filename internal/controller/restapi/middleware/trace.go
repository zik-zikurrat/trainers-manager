package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ctxKey string

const TraceIDKey ctxKey = "trace_id"

func TracingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		traceID := c.Get("X-Trace-ID")
		if traceID == "" {
			traceID = uuid.NewString()
		}
		c.Locals(TraceIDKey, traceID)

		c.Set("X-Trace-ID", traceID)
		return c.Next()
	}
}
