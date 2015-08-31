package goweb

import ()

type Controller struct {
	AppController interface{}

	Request  *Request
	Response *Response
	Result   Result
}

func NewController(req *Request, resp *Response) *Controller {
	return &Controller{
		Request:  req,
		Response: resp,
	}
}
