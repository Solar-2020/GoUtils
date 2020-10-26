package session

import "github.com/valyala/fasthttp"

type BasicRequest struct {

}

type BasicSession struct {}

func NewBasicSession(httpContext *fasthttp.RequestCtx, request BasicRequest) (*BasicSession, error) {
	s := BasicSession{}
	return &s, nil
}
