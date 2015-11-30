package goweb

import "strings"

func StaticHandler(cfg *Config, ctx Context, req Req) bool {
	url := req.Url()
	for _, p := range cfg.ServeFiles {
		if strings.HasPrefix(url, p) {
			ctx.ServeFile(url[1:])
			return false
		}
	}
	return true
}
