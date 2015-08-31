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

func (c *Controller) Html(html string) Result {
	return &HtmlResult{html}
}

func (c *Controller) Json(o interface{}) Result {
	return &JsonResult{o}
}
