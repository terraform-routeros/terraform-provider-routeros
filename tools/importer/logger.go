package main

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ctxLogger struct{}

func NewLogger(ctx context.Context) (*logrus.Logger, context.Context) {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableQuote:     true,
		DisableTimestamp: true,
	})
	log.SetLevel(logrus.DebugLevel)
	// log.SetReportCaller(true)
	return log, ContextWithLogger(ctx, log)
}

// ContextWithLogger adds logger to context
func ContextWithLogger(ctx context.Context, l *logrus.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext returns logger from context
func LoggerFromContext(ctx context.Context) *logrus.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*logrus.Logger); ok {
		return l
	}
	return logrus.StandardLogger()
}
