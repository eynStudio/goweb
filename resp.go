package goweb

import (
	"net/http"
)

type Resp interface {
	http.ResponseWriter
}

type resp struct {
	http.ResponseWriter
}