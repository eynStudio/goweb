package node

import (
	"encoding/json"
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"net/http"
	"os"
	"strings"
)

type Ctx struct {
	di.Container
	Req
	*Resp
	Scope   M
	isErr   bool
	afters  []Handler
	Handled bool
	urlParts
}

func (p *Ctx) Error(code int) *Ctx {
	p.WriteHeader(code)
	p.isErr = true
	return p
}

func (p *Ctx) Set(k string, v T)   { p.Scope[k] = v }
func (p *Ctx) IsErr() bool         { return p.isErr }
func (p *Ctx) Get(k string) string { return p.Scope.GetStr(k) }

func (p *Ctx) NotFound()           { p.Error(http.StatusNotFound) }
func (p *Ctx) Forbidden()          { p.Error(http.StatusForbidden) }
func (p *Ctx) Redirect(url string) { http.Redirect(p.Resp, p.Request, url, http.StatusFound) }

func (p *Ctx) Json(m T) {
	if p.IsErr() {
		return
	}
	if b, err := json.Marshal(m); err != nil {
		p.Error(http.StatusInternalServerError)
	} else {
		p.Resp.Header().Set("Content-Type", "application/json; charset=utf-8")
		p.Resp.Write(b)
	}
}

func (p *Ctx) ServeFile(cfg *Config) bool {
	url := p.Url()
	for _, path := range cfg.ServeFiles {
		if strings.HasPrefix(url, path) {
			if fi, err := os.Stat(url[1:]); err != nil || fi.IsDir() {
				p.NotFound()
			} else {
				http.ServeFile(p.Resp, p.Request, url[1:])
			}
			return true
		}
	}
	return false
}
