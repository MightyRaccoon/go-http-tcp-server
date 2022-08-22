package logger

import (
	"context"

	"go.uber.org/zap"
)

const loggerKey = "LOGGER"

func SetLogger(ctx context.Context) context.Context {
	l, _ := zap.NewProduction()
	return context.WithValue(ctx, loggerKey, l.Sugar())
}

func Fetch(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(loggerKey).(*zap.SugaredLogger)
}
