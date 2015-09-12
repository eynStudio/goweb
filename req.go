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

// Checks if request content-type is json
func (this *Req) IsJsonContent() bool {
	return strings.Contains(this.Header.Get("Content-Type"),"application/json")
}

// Checks if request accept json resp
func (this *Req) IsAcceptJson() bool {
	return strings.Contains(this.Header.Get("Accept"),"application/json")
}