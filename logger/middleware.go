package logger

import (
	"context"
	"log/slog"
)

type HandlerMiddleware struct {
	next slog.Handler
}

func NewHandlerMiddleware(next slog.Handler) *HandlerMiddleware {
	return &HandlerMiddleware{
		next: next,
	}
}

func (hm *HandlerMiddleware) Enabled(ctx context.Context, level slog.Level) bool {
	return hm.next.Enabled(ctx, level)
}

func (hm *HandlerMiddleware) Handle(ctx context.Context, rec slog.Record) error {
	if val, ok := ctx.Value(httpKey).(*HTTPCtx); ok {
		if len(val.Method) != 0 {
			rec.Add("method", val.Method)
		}
		if len(val.RemoteAddr) != 0 {
			rec.Add("remote_addr", val.RemoteAddr)
		}
		if len(val.Host) != 0 {
			rec.Add("host", val.Host)
		}
		if len(val.RequestURI) != 0 {
			rec.Add("request_uri", val.RequestURI)
		}
		if len(val.Referer) != 0 {
			rec.Add("referer", val.Referer)
		}
		if len(val.UserAgent) != 0 {
			rec.Add("user_agent", val.UserAgent)
		}
	}

	return hm.next.Handle(ctx, rec)
}

func (hm *HandlerMiddleware) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &HandlerMiddleware{
		next: hm.next.WithAttrs(attrs),
	}
}

func (hm *HandlerMiddleware) WithGroup(groupName string) slog.Handler {
	return &HandlerMiddleware{
		next: hm.next.WithGroup(groupName),
	}
}

type HTTPCtx struct {
	Method     string
	RemoteAddr string
	RequestURI string
	Host       string
	Referer    string
	UserAgent  string
}

type keyType int

const httpKey keyType = 0
