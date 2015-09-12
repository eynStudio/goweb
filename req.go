package goweb

import (
	"strings"
	"net/http"
)

type Req struct{
	*http.Request
}

func (this *Req) Url() string{
	return this.URL.Path 	
}

func (this *Req) Method() string{
	return strings.ToLower(this.Request.Method)
}