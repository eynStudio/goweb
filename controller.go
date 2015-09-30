package goweb

import (
	"reflect"
)

type Controller struct {
	Ctx    Context `di`
	Req    Req     `di`
	Params Values  `di`
}

type CtrlInfo struct {
	Name    string
	Auth	bool
	Type    reflect.Type
	Methods map[string]CtrlAction
}

type CtrlAction struct {
	Name string
	Args []reflect.Type
}

func NewCtrlInfo(name string, t reflect.Type,needAuth bool) *CtrlInfo {
	return &CtrlInfo{
		Name:    name,
		Type:    t,
		Auth:needAuth,
		Methods: make(map[string]CtrlAction),
	}
}
