package goweb

import (
	"encoding/json"
)

type Result interface {
	Apply(req *Request, resp *Response)
}

type ErrorResult struct {
	Error error
}

func (this ErrorResult) Apply(req *Request, resp *Response) {
	resp.Header("Content-Type", "text/html; charset=utf-8")
	resp.Out.Write([]byte(this.Error.Error()))
}

type HtmlResult struct {
	Html string
}

func (this HtmlResult) Apply(req *Request, resp *Response) {
	resp.Header("Content-Type", "text/html; charset=utf-8")
	resp.Out.Write([]byte(this.Html))
}

type JsonResult struct {
	Data interface{}
}

func (this JsonResult) Apply(req *Request, resp *Response) {
	resp.Header("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(this.Data)
	if err != nil {
		ErrorResult{Error: err}.Apply(req, resp)
		return
	}

	resp.Out.Write(b)
}
