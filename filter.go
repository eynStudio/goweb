package goweb

import (
	"os"
	"strings"
)

func StaticHandler(cfg *Config, ctx Context, req Req) bool {
	url := req.Url()
	for _, p := range cfg.ServeFiles {
		if strings.HasPrefix(url, p) {
			if fi, err := os.Stat(url[1:]); err != nil || fi.IsDir() {
				ctx.NotFound()
			} else {
				ctx.ServeFile(url[1:])
			}
			return false
		}
	}
	return true
}
