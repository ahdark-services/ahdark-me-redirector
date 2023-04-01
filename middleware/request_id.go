package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), "request-id")
		defer span.End()
		c.Request = c.Request.WithContext(ctx)

		requestId := uuid.New().String()

		c.Set("request_id", requestId)
		c.Header("X-Request-Id", requestId)

		span.SetAttributes(attribute.String("request_id", requestId))

		c.Next()

		logrus.WithContext(ctx).WithField("request_id", requestId).
			Debugf("request to %s completed with status %d", c.Request.URL.Path, c.Writer.Status())
	}
}
