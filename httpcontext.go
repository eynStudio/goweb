package goweb

import (
	"net/http"
)

type HttpContext struct {
	Req    *http.Request
	Resp   http.ResponseWriter
	App    *App
	Route  *Route
	Params map[string]string
	Result Result
}

func NewHttpContext(r *http.Request, rw http.ResponseWriter) (ctx *HttpContext) {
	ctx = &HttpContext{Req: r, Resp: rw, Params: make(map[string]string)}

	return
}

func (this *HttpContext) Header(key, val string) {
	this.Resp.Header().Set(key, val)
}
