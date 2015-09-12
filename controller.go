package goweb

import (
	"reflect"
)

type Controller interface {
}

type CtrlInfo struct {
	Name    string
	Type    reflect.Type
	Methods map[string]CtrlAction
}
type CtrlAction struct {
	Name string
	Args []reflect.Type
}

func NewCtrlInfo(name string, t reflect.Type) *CtrlInfo {
	return &CtrlInfo{
		Name:    name,
		Type:    t,
		Methods: make(map[string]CtrlAction),
	}
}
