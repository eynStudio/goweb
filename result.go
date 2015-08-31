package goweb

import ()

type Result interface {
	Apply(req *Request, resp *Response)
}
