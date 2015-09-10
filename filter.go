package goweb

import (
	"path"
	"strings"
)

func StaticHandler(ctx Context) bool {
	url := path.Clean(ctx.(*context).Req.URL.Path)
	if strings.HasPrefix(url, "/static") {
		ctx.ServeFile(url[1:])
		return false
	}
	return true
}
