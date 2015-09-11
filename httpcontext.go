package goweb

import (
	"github.com/eynstudio/gobreak/di"
	"net/http"
)

type Context interface {
	di.Container
	ServeFile(path string)
	Tmpl(tpl string, o interface{})
	Html(html string)
	Json(o interface{})
	Forbidden()
}
type context struct {
	di.Container
	Req      *http.Request
	Resp     http.ResponseWriter
	handlers []Handler
	Params   map[string]string
	Result   Result
}

func (this *context) Header(key, val string) {
	this.Resp.Header().Set(key, val)
}

func (this *context) ServeFile(path string) {
	http.ServeFile(this.Resp, this.Req, path)
}

func (this *context) Tmpl(tpl string, o interface{}) {
	this.Result = &TemplateResult{tpl, o}
}

func (this *context) Html(html string) {
	this.Result = &HtmlResult{html}
}

func (this *context) Json(o interface{}) {
	this.Result = &JsonResult{o}
}

func (this *context) Forbidden() {
	this.Result = ResultForbidden
}
func (this *context) exec() {
	for _, it := range this.handlers {
		next, err := this.Invoke(it)
		if err != nil {
			panic(err)
		}

		if next != nil && !next[0].Bool() {
			return
		}
	}
}
