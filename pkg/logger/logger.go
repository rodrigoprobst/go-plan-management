package logger

import (
	"context"
	"errors"

	"go.uber.org/zap"
)

type contextKey struct{}

func NewLogger(app, env string) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stdout"}
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.InitialFields = map[string]interface{}{
		"app":         app,
		"environment": env,
	}

	return cfg.Build()
}

func NewNop() *zap.Logger {
	return zap.NewNop()
}

func AddToContext(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, contextKey{}, l)
}

func FromContext(ctx context.Context) *zap.Logger {
	l, ok := ctx.Value(contextKey{}).(*zap.Logger)
	if !ok {
		panic(errors.New("logger not found"))
	}

	return l
}
