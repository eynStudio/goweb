package goweb

import (
	"reflect"
)

type Handler interface{}

func checkHandler(handler Handler) {
	if reflect.TypeOf(handler).Kind() != reflect.Func {
		panic("handler must be a func")
	}
}
