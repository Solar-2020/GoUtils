package errorWorker

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type ErrorWorker interface {
	ServeJSONError(ctx *fasthttp.RequestCtx, serveError error)
	ServeFatalError(ctx *fasthttp.RequestCtx)
	NewError(httpCode int, responseError error, fullError error) (err error)
}

type errorWorker struct {
	defaultCode int
}

func NewErrorWorker() ErrorWorker {
	return &errorWorker{defaultCode: fasthttp.StatusBadRequest}
}

type ServeError struct {
	Error interface{} `json:"error"`
}

func (ew *errorWorker) NewError(httpCode int, responseError error, fullError error) (err error) {
	return ResponseError{
		httpCode:      httpCode,
		responseError: responseError,
		fullError:     fullError,
	}
}

func (ew *errorWorker) ServeJSONError(ctx *fasthttp.RequestCtx, serveError error) {
	if responseError, ok := serveError.(*ResponseError); ok {
		ctx.SetUserValue("error", responseError.fullError)
		ew.serveJSONError(ctx, responseError.httpCode, responseError.responseError.Error())
		return
	}
	ctx.SetUserValue("error", serveError)
	ew.serveJSONError(ctx, ew.defaultCode, serveError.Error())

	return
}

func (ew *errorWorker) serveJSONError(ctx *fasthttp.RequestCtx, statusCode int, err interface{}) {
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(statusCode)

	errorStruct := ServeError{Error: err}

	body, marshalErr := json.Marshal(errorStruct)
	if marshalErr != nil {
		ew.sendInternalError(ctx)
		return
	}

	ctx.SetBody(body)
}

func (ew *errorWorker) sendInternalError(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetStatusCode(fasthttp.StatusInternalServerError)
	return
}

func (ew *errorWorker) ServeFatalError(ctx *fasthttp.RequestCtx) {
	ew.sendInternalError(ctx)
}
