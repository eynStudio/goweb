package node

import (
	"net/http"
	"strings"
)

type Handler func(c *Ctx)

type Interceptor struct {
	Before Handler
	After  Handler
}

type Req struct {
	*http.Request
}

func (p *Req) Url() string {
	return p.URL.Path
}

func (p *Req) Method() string {
	return strings.ToLower(p.Request.Method)
}

func (p *Req) IsJsonContent() bool {
	return strings.Contains(p.Header.Get("Content-Type"), "application/json")
}

func (p *Req) IsAcceptJson() bool {
	return strings.Contains(p.Header.Get("Accept"), "application/json")
}

type Resp struct {
	http.ResponseWriter
}
