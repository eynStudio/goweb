package goweb

import (
	"encoding/json"
	"net/http"
)

var (
	ResultForbidden ErrorResult = ErrorResult{"Forbidden", http.StatusForbidden}
	ResultNotFound  ErrorResult = ErrorResult{"NotFound", http.StatusNotFound}
)

type Result interface {
	Apply(ctx *HttpContext)
}

type ErrorResult struct {
	Msg      string
	HttpCode int
}

func (this ErrorResult) Apply(ctx *HttpContext) {
	http.Error(ctx.Resp, this.Msg, this.HttpCode)
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
		ErrorResult{"InternalServerError", http.StatusInternalServerError}.Apply(ctx)
		return
	}

	ctx.Resp.Write(b)
}
