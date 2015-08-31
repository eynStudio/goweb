package goweb

import (
	"net/http"
)

type Request struct {
	*http.Request
}

type Response struct {
	Status      int
	ContentType string

	Out http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{Out: w}
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		Request: r,
	}
}

// Header sets response header item string via given key.
func (this *Response) Header(key, val string) {
	this.Out.Header().Set(key, val)
}
