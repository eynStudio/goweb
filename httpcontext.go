package goweb

import (
	"net/http"
	"path"
	"strings"
)

type HttpContext struct {
	Req    *http.Request
	Resp   http.ResponseWriter
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

func StaticFilter(ctx *HttpContext, fc []Filter) {
	url := path.Clean(ctx.Req.URL.Path)
	if strings.HasPrefix(url, "/static") {
		http.ServeFile(ctx.Resp, ctx.Req, url[1:])
	} else {
		fc[0](ctx, fc[1:])
	}
}
