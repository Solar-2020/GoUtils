package context

import (
	"encoding/json"
	"fmt"
	"github.com/Solar-2020/GoUtils/context/session"
	"github.com/valyala/fasthttp"
)

// With authorisation
// Data for request
type Context struct {
	*fasthttp.RequestCtx
	Session *session.Session
}

type SpecialRequest struct {
	session.RequestWithAuth
}

func NewContext(httpCtx *fasthttp.RequestCtx) (ctx Context, err error) {
	req := SpecialRequest{}
	body := httpCtx.Request.Body()
	if body != nil && len(body) > 0{
		err = json.Unmarshal(body, &req)
		if err != nil {
			return Context{}, err
		}
	}
	ctx = Context{
		RequestCtx: httpCtx,
		Session:    nil,
	}
	err = ctx.Inflate(httpCtx, &req)

	return ctx, err
}

func NewEmptyContext(httpCtx *fasthttp.RequestCtx) (ctx Context, err error) {
	ctx = Context{
		RequestCtx: httpCtx,
		Session:    nil,
	}
	s, err := session.NewEmptySession(httpCtx)
	ctx.Session = s
	return
}

func NewMockContext(httpCtx *fasthttp.RequestCtx) (ctx Context, err error) {
	req := SpecialRequest{}
	body := httpCtx.Request.Body()
	if body != nil && len(body) > 0{
		err = json.Unmarshal(body, &req)
		if err != nil {
			return Context{}, err
		}
	}
	ctx = Context{
		RequestCtx: httpCtx,
		Session:    nil,
	}
	s, err := session.NewMockSession(httpCtx, req.RequestWithAuth)
	if err != nil {
		return
	}
	s.God = true
	ctx.Session = s
	return ctx, err
}

func (c *Context) Inflate(httpCtx *fasthttp.RequestCtx, req *SpecialRequest) error {
	session, err := session.NewSession(httpCtx, req.RequestWithAuth)
	c.Session = session
	return err
}

func (c *Context) InflateOpen(httpCtx *fasthttp.RequestCtx, req *SpecialRequest) error {
	s, err := session.NewBasicSession(httpCtx, req.RequestWithAuth.BasicRequest)
	if err != nil {
		return err
	}
	if s == nil {
		return fmt.Errorf("nil session")
	}
	c.Session = &session.Session{
		BasicSession: *s,
	}
	return err
}

func (c *Context) GetUid() int {
	return c.Session.Uid
}

func (c *Context) GetLogin() string {
	return c.Session.Login
}


func NewOpenContext(httpCtx *fasthttp.RequestCtx) (Context, error) {
	req := SpecialRequest{}
	err := json.Unmarshal(httpCtx.Request.Body(), &req)
	if err != nil {
		return Context{}, err
	}
	ctx := Context{
		RequestCtx: httpCtx,
		Session:    nil,
	}
	err = ctx.InflateOpen(httpCtx, &req)

	return ctx, err
}





