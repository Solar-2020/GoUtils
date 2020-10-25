package context

import (
	"encoding/json"
	"fmt"
	"github.com/Solar-2020/GoUtils/context/session"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestNewContext(t *testing.T) {
	req := SpecialRequest{
		RequestWithAuth: session.RequestWithAuth{
			Uid: 123,
		},
	}

	body, _ := json.Marshal(req)
	req2 := fasthttp.Request{}
	req2.AppendBody(body)
	ctx := fasthttp.RequestCtx{
		Request:  req2,
		Response: fasthttp.Response{},
	}
	c, _ := NewContext(&ctx)
	fmt.Println(c)
}
