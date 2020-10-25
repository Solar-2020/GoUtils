package log

import (
	"context"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

const (
	ctxLoggerKey = "X-Logger"
	timeFormat = "02-Jan-2006 15:04:05"
)


func Println(ctx_ interface{}, s ...interface{}) {
	ctx, ok := ctx_.(*fasthttp.RequestCtx)
	if !ok {
		return
	}
	logger := Get(ctx)
	if logger == nil {
		return
	}
	logger.Println(ctx, s...)
}

func Printf(ctx_ interface{}, format string, s ...interface{}) {
	ctx, ok := ctx_.(*fasthttp.RequestCtx)
	if !ok {
		return
	}
	logger := Get(ctx)
	if logger == nil {
		return
	}
	logger.Printf(ctx, format, s...)
}


type Log interface {
	Println(ctx context.Context, s ...interface{})
	Printf(ctx context.Context, format string, s ...interface{})
}

func Get(ctx *fasthttp.RequestCtx) Log {
	return ctx.Value(ctxLoggerKey).(Log)
}

func Set(ctx *fasthttp.RequestCtx, logger Log) {
	ctx.SetUserValue(ctxLoggerKey, logger)
}

type PassportContext interface {
	GetUid() int
	GetLogin() string
}

type log struct {}

func NewLog() log {
	return log{}
}

func (l *log) Println(ctx context.Context, s ...interface{}) {
	fmt.Printf(l.preformat(ctx), fmt.Sprint(s...))
}

func (l *log) Printf(ctx context.Context, format string, s ...interface{}) {
	end := fmt.Sprintf(format, s...)
	fmt.Printf(l.preformat(ctx), end)
}

func (l *log) preformat(ctx context.Context) string {
	if newctx, ok := ctx.(PassportContext); ok {
		return fmt.Sprintf(
			"[%s uid=%d login=%s] %s\n",
			time.Now().Format(timeFormat), newctx.GetUid(), newctx.GetLogin(), "%s",
		)
	}
	return fmt.Sprintf(
		"[%s] %s\n",
		time.Now().Format(timeFormat), "%s",
	)
}
