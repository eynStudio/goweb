package goweb

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	ResultForbidden ErrorResult  = ErrorResult{"Forbidden", http.StatusForbidden}
	ResultNotFound  ErrorResult  = ErrorResult{"NotFound", http.StatusNotFound}
	ResulOK         StatusResult = StatusResult{http.StatusOK}
)

type Result interface {
	Apply(ctx *context)
}

type ErrorResult struct {
	Msg      string
	HttpCode int
}

func (p ErrorResult) Apply(ctx *context) {
	http.Error(ctx.Resp, p.Msg, p.HttpCode)
}

type StatusResult struct {
	HttpCode int
}

func (p StatusResult) Apply(ctx *context) {
	ctx.Resp.WriteHeader(p.HttpCode)
}

type TemplateResult struct {
	Tpl  string
	Data interface{}
}

func (p TemplateResult) Apply(ctx *context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	if err := MyTemplates.Execute(ctx.Resp, p.Tpl, p.Data); err != nil {
		log.Println(err)
	}
}

type HtmlResult struct {
	Html string
}

func (p HtmlResult) Apply(ctx *context) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	ctx.Resp.Write([]byte(p.Html))
}

type JsonResult struct {
	Data interface{}
}

func (p JsonResult) Apply(ctx *context) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(p.Data)
	if err != nil {
		ErrorResult{"InternalServerError", http.StatusInternalServerError}.Apply(ctx)
		return
	}

	ctx.Resp.Write(b)
}
