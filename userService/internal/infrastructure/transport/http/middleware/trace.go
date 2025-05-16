package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const TraceIDKey = "traceId"

func TraceMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-Trace-Id")
		if traceID == "" {
			traceID = uuid.New().String()
		}
		ctx := context.WithValue(c.Request.Context(), TraceIDKey, traceID)
		requestLogger := logger.With(zap.String("traceId", traceID))
		c.Set("logger", requestLogger)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
