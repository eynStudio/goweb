package goweb

import (
	"encoding/json"
)

type Result interface {
	Apply(ctx *HttpContext)
}

type ErrorResult struct {
	Error error
}

func (this ErrorResult) Apply(ctx *HttpContext) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Resp.Write([]byte(this.Error.Error()))
}

type TemplateResult struct {
	Tpl  string
	Data interface{}
}

func (this TemplateResult) Apply(ctx *HttpContext) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	MyTemplates.Execute(ctx.Resp, this.Tpl, this.Data)
}

type HtmlResult struct {
	Html string
}

func (this HtmlResult) Apply(ctx *HttpContext) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Resp.Write([]byte(this.Html))
}

type JsonResult struct {
	Data interface{}
}

func (this JsonResult) Apply(ctx *HttpContext) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(this.Data)
	if err != nil {
		ErrorResult{Error: err}.Apply(ctx)
		return
	}

	ctx.Resp.Write(b)
}
