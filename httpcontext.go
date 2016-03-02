package goweb

import (
	"net/http"

	"github.com/eynstudio/gobreak/di"
)

type Context interface {
	di.Container
	Header(key, val string)
	ServeFile(path string)
	Tmpl(tpl string, o interface{})
	Html(html string)
	Json(o interface{})
	Redirect(url string)
	Ok()
	NotFound()
	Forbidden()
}
type context struct {
	di.Container
	Req
	Resp
	handlers []Handler
	Params   map[string]string
	Result   Result
}

func (p *context) Header(key, val string) {
	p.Resp.Header().Set(key, val)
}

func (p *context) Ok() {
	//	this.Resp.WriteHeader(http.StatusOK)
	p.Result = ResulOK
}

func (p *context) Redirect(url string) {
	http.Redirect(p.Resp, p.Req.Request, url, http.StatusFound)
}

func (p *context) NotFound() {
	p.Result = ResultNotFound
	//	this.Resp.WriteHeader(http.StatusNotFound)
}
func (p *context) ServeFile(path string) {
	http.ServeFile(p.Resp, p.Req.Request, path)
}

func (p *context) Tmpl(tpl string, o interface{}) {
	p.Result = &TemplateResult{tpl, o}
}

func (p *context) Html(html string) {
	p.Result = &HtmlResult{html}
}

func (p *context) Json(o interface{}) {
	p.Result = &JsonResult{o}
}

func (p *context) Forbidden() {
	p.Result = ResultForbidden
}
func (p *context) exec() {
	for _, it := range p.handlers {
		next, err := p.Invoke(it)
		if err != nil {
			panic(err)
		}

		if next != nil && !next[0].Bool() {
			return
		}
	}
}
