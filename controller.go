package goweb

import (
	"reflect"
)

type Controller interface {
	SetCtx(ctx *context)
}

type BaseController struct {
	Ctx *context `di`
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

func NewControllerInfo(name string, t reflect.Type) *ControllerInfo {
	return &ControllerInfo{
		Name:    name,
		Type:    t,
		Methods: make(map[string]ControllerMethod),
	}
}

func (this *BaseController) SetCtx(ctx *context) {
	this.Ctx = ctx
}

func (this *BaseController) Html(html string) Result {
	return &HtmlResult{html}
}

func (this *BaseController) Json(o interface{}) Result {
	return &JsonResult{o}
}
