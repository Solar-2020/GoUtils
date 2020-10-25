package errorWorker

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ErrorWorker interface {
	ServeJSONError(ctx *fasthttp.RequestCtx, serveError error) (err error)
	ServeFatalError(ctx *fasthttp.RequestCtx)
}

type errorWorker struct {
}

func NewErrorWorker() ErrorWorker {
	return &errorWorker{}
}

type ServeError struct {
	Error string `json:"error"`
}

func (ew *errorWorker) ServeJSONError(ctx *fasthttp.RequestCtx, serveError error) (err error) {
	sendError := ServeError{
		Error: serveError.Error(),
	}

	body, err := json.Marshal(sendError)
	if err != nil {
		return
	}

	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusBadRequest)
	ctx.SetBody(body)

	return
}

func (ew *errorWorker) ServeFatalError(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetStatusCode(fasthttp.StatusInternalServerError)
	return
}
