package goweb

import (
	"net/http"
	"path"
	"strings"
)

type Filter func(ctx *context, filterChain []Filter)

var Filters = []Filter{
	StaticFilter,
	RouterFilter,
	ControllerFilter,
	func(*context, []Filter) {},
}

func StaticFilter(ctx *context, fc []Filter) {
	url := path.Clean(ctx.Req.URL.Path)
	if strings.HasPrefix(url, "/static") {
		http.ServeFile(ctx.Resp, ctx.Req, url[1:])
	} else {
		fc[0](ctx, fc[1:])
	}
}
