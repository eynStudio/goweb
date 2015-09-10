package goweb

import (
	"net/http"
	"path"
	"strings"
)

func StaticHandler(ctx Context) bool {
	url := path.Clean(ctx.(*context).Req.URL.Path)
	if strings.HasPrefix(url, "/static") {
		http.ServeFile(ctx.(*context).Resp, ctx.(*context).Req, url[1:])
		return false
	}
	return true
}
