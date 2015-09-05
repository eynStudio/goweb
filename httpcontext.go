package goweb

import (
	"net/http"
)

type HttpContext struct {
	Req    *http.Request
	Resp   http.ResponseWriter
	Result Result
}

func NewHttpContext(r *http.Request, rw http.ResponseWriter) *HttpContext {
	return &HttpContext{Req: r, Resp: rw}
}

func (this *HttpContext) Header(key, val string) {
	this.Resp.Header().Set(key, val)
}
