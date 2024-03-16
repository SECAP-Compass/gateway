package server

import (
	"context"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func correlationIdMiddleware(c *fiber.Ctx) error {
	correlationId := c.Get("X-Correlation-ID")
	if correlationId == "" {
		correlationId = uuid.NewString()
	}
	c.Set("X-Correlation-ID", correlationId)
	_, span := StartSpan(c.UserContext())
	span.SetAttributes(attribute.String("correlation_id", correlationId))

	return c.Next()
}

var Tracer = otel.GetTracerProvider().Tracer("secap-compass")

func StartSpan(ctx context.Context) (context.Context, trace.Span) {
	pc, _, _, _ := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	return Tracer.Start(ctx, details.Name())
}
