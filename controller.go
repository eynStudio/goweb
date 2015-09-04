package goweb

import (
	"reflect"
)

type Controller struct {
	AppController interface{}

	Request  *Request
	Response *Response
	Result   Result
}

type ControllerInfo struct {
	Name    string
	Type    reflect.Type
	Methods []*ControllerMethod
}
type ControllerMethod struct {
	Name string
	Args []*ControllerMethodArg
}
type ControllerMethodArg struct {
	Name string
	Type reflect.Type
}

func NewController(req *Request, resp *Response) *Controller {
	return &Controller{
		Request:  req,
		Response: resp,
	}
}

func (c *Controller) Html(html string) Result {
	return &HtmlResult{html}
}

func (c *Controller) Json(o interface{}) Result {
	return &JsonResult{o}
}
