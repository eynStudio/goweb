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
	Methods map[string]ControllerMethod
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

func NewControllerInfo(name string, t reflect.Type) *ControllerInfo {
	return &ControllerInfo{
		Name:    name,
		Type:    t,
		Methods: make(map[string]ControllerMethod),
	}
}
func (c *Controller) Html(html string) Result {
	return &HtmlResult{html}
}

func (c *Controller) Json(o interface{}) Result {
	return &JsonResult{o}
}
