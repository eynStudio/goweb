package goweb

import (
	"github.com/eynStudio/gobreak/di"
	"net/http"
)

type Context interface {
	di.Container
	Next()
}
type context struct {
	di.Container
	Req    *http.Request
	Resp   http.ResponseWriter
	App    *App
	Route  *Route
	Params map[string]string
	Result Result
}

func (this *context) Header(key, val string) {
	this.Resp.Header().Set(key, val)
}

func (this *context) Next() {}
