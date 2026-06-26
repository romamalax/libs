package logger

import (
	"context"
	"log/slog"
	"os"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger(level string) *Logger {
	l := &Logger{}

	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "error":
		logLevel = slog.LevelError
	case "warn":
		logLevel = slog.LevelWarn
	default:
		logLevel = slog.LevelDebug
	}

	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	handler := NewHandlerMiddleware(jsonHandler)

	l.log = slog.New(handler)
	slog.SetDefault(l.log)

	return l
}

func (l *Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.log.Error(msg, args...)
	os.Exit(1)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.log.DebugContext(ctx, msg, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.log.InfoContext(ctx, msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.log.ErrorContext(ctx, msg, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.log.WarnContext(ctx, msg, args...)
}

func (l *Logger) WithHTTPContext(ctx context.Context, method, remoteAddr, requestURI string) context.Context {
	httpCtx := &HTTPCtx{
		Method:     method,
		RemoteAddr: remoteAddr,
		RequestURI: requestURI,
	}
	return context.WithValue(ctx, httpKey, httpCtx)
}

func (l *Logger) WithHTTPContextFull(ctx context.Context, method, remoteAddr, host, requestURI, referer, userAgent string) context.Context {
	httpCtx := &HTTPCtx{
		Method:     method,
		RemoteAddr: remoteAddr,
		Host:       host,
		RequestURI: requestURI,
		Referer:    referer,
		UserAgent:  userAgent,
	}
	return context.WithValue(ctx, httpKey, httpCtx)
}
