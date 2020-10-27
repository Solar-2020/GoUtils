package session

import (
	"fmt"
	"github.com/Solar-2020/GoUtils/log"
	"github.com/valyala/fasthttp"
	"strconv"
)

const (
	UidHeaderKey = "X-UID"
	EmailHeaderKey = "X-Login"
	cookieTokenKey = "SessionToken"
)

type AuthClient interface {
	ByCookie(cookie string, headers map[string]string) (int, error)
}

var (
	authClient AuthClient
)

func RegisterAuthService(concreteAuthClient AuthClient) {
	authClient = concreteAuthClient
}

type Session struct {
	BasicSession
	Uid int
	Login  string
}

type RequestWithAuth struct {
	BasicRequest
	Uid int `json:"uid"`
}

func NewSession(httpCtx *fasthttp.RequestCtx, request RequestWithAuth) (*Session, error) {
	s := Session{}
	basicS, err := NewBasicSession(httpCtx, request.BasicRequest)
	if err != nil {
		return nil, err
	}
	s.BasicSession = *basicS

	err = s.Authorise(httpCtx, request)

	return &s, err
}

func (s *Session) Authorise(ctx *fasthttp.RequestCtx, request RequestWithAuth) error{
	return s.serviceAuthorise(ctx, request)
	//return s.mockAuthorise(ctx, request)
}

func (s *Session) serviceAuthorise(ctx *fasthttp.RequestCtx, request RequestWithAuth) error{
	err := s.mockAuthorise(ctx, request)
	sessionCookieSrc := ctx.Request.Header.Cookie(cookieTokenKey)
	if sessionCookieSrc == nil {
		if s.Uid == 0  {
			return fmt.Errorf("token required")
		}
		log.Println(ctx, "auth with header")
		return nil
	}
	sessionCookie := string(sessionCookieSrc)
	fmt.Println(sessionCookie)
	if authClient == nil {
		if s.Uid == 0  {
			return fmt.Errorf("nil auth service")
		}
		log.Println(ctx, "auth service needed")
		return nil
	}
	uid, err := authClient.ByCookie(sessionCookie, nil)
	if err != nil {
		return err
	}
	if uid != 0 {
		s.Uid = uid
	}
	return nil
}

func (s *Session) mockAuthorise(ctx *fasthttp.RequestCtx, request RequestWithAuth) error {
	s.Uid = request.Uid

	if s.Uid == 0 {
		if headerUid, err := s.uidFromHeader(ctx); err == nil {
			log.Println(ctx, "Uid переопределен из заголовка: ", headerUid)
			s.Uid = headerUid
		}
	}

	uidToEmail := func(uid int) string {
		return fmt.Sprintf("email_uid_%d@solar.ru", uid)
	}

	s.Login = uidToEmail(s.Uid)
	if headerLogin, err := s.emailFromHead(ctx); err == nil {
		log.Println(ctx, "Email переопределен из заголовка: ", headerLogin)
		s.Login = headerLogin
	}

	return nil
}

func (s *Session) emailFromHead(ctx *fasthttp.RequestCtx) (string, error) {
	b := ctx.Request.Header.Peek(EmailHeaderKey)
	if b == nil {
		return "", fmt.Errorf("missed")
	}
	return string(b), nil
}

func (s *Session) uidFromHeader(ctx *fasthttp.RequestCtx) (int, error) {
	b := ctx.Request.Header.Peek(UidHeaderKey)
	if b == nil {
		return 0, fmt.Errorf("missed")
	}
	header := string(b)
	return strconv.Atoi(header)
}

const (
	ctxKey = "X-Session-Key"
)

func NewToCtx(ctx *fasthttp.RequestCtx, session interface{}){
	ctx.SetUserValue(ctxKey, session)
}

func GetAnon(ctx *fasthttp.RequestCtx) (s BasicSession, err error) {
	res := ctx.Value(ctxKey)
	if res == nil {
		err = fmt.Errorf("No session")
		return
	}
	s = res.(BasicSession)
	return
}

func GetAuthorized(ctx *fasthttp.RequestCtx) (s Session, err error) {
	res := ctx.Value(ctxKey)
	if res == nil {
		err = fmt.Errorf("No session")
		return
	}
	s = res.(Session)
	return
}

// TODO: go to auth service