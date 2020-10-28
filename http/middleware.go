package http

import (
	"github.com/Solar-2020/GoUtils/context"
	"github.com/Solar-2020/GoUtils/log"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type CleanHandler func(ctx context.Context)

type Middleware interface {
	CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler
	Log(next fasthttp.RequestHandler) fasthttp.RequestHandler
	Auth(next CleanHandler) fasthttp.RequestHandler
	InternalAuth(next CleanHandler) fasthttp.RequestHandler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return middleware{}
}

var (
	corsAllowHeaders     = "authorization"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func (m middleware) CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		//ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		//ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		//ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		//ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}

func (m middleware) Log(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logger := log.NewLog()
		log.Set(ctx, &logger)
		logger.Println(ctx, "Start new request: ", ctx.Request.URI())
		logger.Println(ctx, ctx.Request.String())

		defer func(begin time.Time) {
			logger.Printf(
				ctx,
				"End: %s, status: %d, time: %d ms",
				ctx.Request.URI().String(),
				ctx.Response.StatusCode(),
				time.Since(begin).Milliseconds(),
			)
		}(time.Now())

		next(ctx)
	}
}

func (m middleware) Auth(next CleanHandler) fasthttp.RequestHandler {
	return func(httpCtx *fasthttp.RequestCtx) {
		ctx, err := context.NewContext(httpCtx)
		if err != nil {
			httpCtx.Response.SetStatusCode(http.StatusForbidden)
			return
		}
		log.Println(ctx, "Auth successful")
		next(ctx)
	}
}

func (m middleware) InternalAuth(next CleanHandler) fasthttp.RequestHandler {
	return func(httpCtx *fasthttp.RequestCtx) {
		ctx, err := context.NewMockContext(httpCtx)
		if err != nil {
			log.Println(httpCtx, err)
			httpCtx.Response.SetStatusCode(http.StatusInternalServerError)
			return
		}
		next(ctx)
	}
}

func NewLogCorsChain(middleware Middleware) func(func(ctx *fasthttp.RequestCtx)) fasthttp.RequestHandler {
	return func(target func(ctx *fasthttp.RequestCtx)) fasthttp.RequestHandler {
		return middleware.Log(middleware.CORS(target))
	}
}

func ClientsideChain(middleware Middleware)  func(CleanHandler) fasthttp.RequestHandler {
	return func(target CleanHandler) fasthttp.RequestHandler {
		return middleware.Log(middleware.CORS(middleware.Auth(target)))
	}
}

func ServersideChain(middleware Middleware)  func(CleanHandler) fasthttp.RequestHandler {
	return func(target CleanHandler) fasthttp.RequestHandler {
		return middleware.Log(middleware.CORS(middleware.InternalAuth(target)))
	}
}