package http

import (
	"fmt"
	"github.com/Solar-2020/GoUtils/http/errorWorker"
	"github.com/valyala/fasthttp"
	"runtime/debug"
)

func PanicHandler(ctx *fasthttp.RequestCtx, err interface{}) {
	fmt.Printf("Request falied with panic: %s, error: %v\nTrace:\n", string(ctx.Request.RequestURI()), err)
	fmt.Println(string(debug.Stack()))
	errorWorker.NewErrorWorker().ServeFatalError(ctx)
}

func HealthCheckHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("X-Health-Check", "43234")	// random number
}