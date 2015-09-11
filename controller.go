package goweb

import (
	"reflect"
)

type Controller interface {
}

type ControllerInfo struct {
	Name    string
	Type    reflect.Type
	Methods map[string]ControllerMethod
}
type ControllerMethod struct {
	Name string
	Args []reflect.Type
}

func NewControllerInfo(name string, t reflect.Type) *ControllerInfo {
	return &ControllerInfo{
		Name:    name,
		Type:    t,
		Methods: make(map[string]ControllerMethod),
	}
}
