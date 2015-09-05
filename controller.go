package goweb

import (
	"reflect"
)

type Controller struct {
	Ctx    *HttpContext
	Result Result
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

func NewController(ctx *HttpContext) *Controller {
	return &Controller{Ctx: ctx}
}

func (c *Controller) Html(html string) Result {
	return &HtmlResult{html}
}

func (c *Controller) Json(o interface{}) Result {
	return &JsonResult{o}
}
