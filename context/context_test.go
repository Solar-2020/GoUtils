package context

import (
	"encoding/json"
	"fmt"
	"github.com/Solar-2020/GoUtils/context/session"
	"github.com/valyala/fasthttp"
	"net/url"
	"path"
	"testing"
)

func TestNewContext(t *testing.T) {
	url, err := url.Parse("http://localhost:3000/auth/")
	if err != nil {
		return
	}
	url.Path = path.Join(url.Path, "/authorization/signup")
	res := url.String()
	fmt.Println(res)


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
