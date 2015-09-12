package goweb

import (
	"strings"
)

func StaticHandler(ctx Context,req Req) bool {
	url :=req.Url()
	if strings.HasPrefix(url, "/static") {
		ctx.ServeFile(url[1:])
		return false
	}
	return true
}
