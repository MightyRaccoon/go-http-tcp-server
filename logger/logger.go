package logger

import (
	"context"
	"lowlevelserver/utils"

	"go.uber.org/zap"
)


func SetLogger(ctx context.Context) context.Context {
	l, _ := zap.NewProduction()
	return context.WithValue(ctx, utils.LoggerKey, l.Sugar())
}

func Fetch(ctx context.Context) *zap.SugaredLogger {
	return ctx.Value(utils.LoggerKey).(*zap.SugaredLogger)
}
