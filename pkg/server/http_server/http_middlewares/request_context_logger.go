package http_middlewares

import (
	"github.com/rodrigoprobst/go-plan-management/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func LoggerMiddleware(l *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		span, _ := tracer.SpanFromContext(c.Request.Context())
		traceID := span.Context().TraceID()
		spanID := span.Context().SpanID()

		log := l.With(
			zap.Uint64("trace_id", traceID),
			zap.Uint64("span_id", spanID),
		)

		ctx := logger.AddToContext(c.Request.Context(), log)
		c.Request = c.Request.WithContext(ctx)

		request := map[string]interface{}{
			"uri":        c.Request.RequestURI,
			"headers":    c.Request.Header,
			"method":     c.Request.Method,
			"user_agent": c.Request.UserAgent(),
		}

		log.Info("incoming request",
			zap.Any("request", request),
		)

		c.Next()

		response := map[string]interface{}{
			"status": c.Writer.Status(),
		}

		if len(c.Errors) > 0 {
			response["errors"] = c.Errors.String()
		}

		log.Info("finish request",
			zap.Any("response", response),
		)

	}
}
