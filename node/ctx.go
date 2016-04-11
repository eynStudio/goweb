package node

import (
	"encoding/json"
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
	"net/http"
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

func (p *Ctx) Error(code int, msg string) *Ctx {
	http.Error(p.Resp, msg, code)
	p.isErr = true
	return p
}

func (p *Ctx) Set(k string, v T)   { p.Scope[k] = v }
func (p *Ctx) IsErr() bool         { return p.isErr }
func (p *Ctx) Get(k string) string { return p.Scope.GetStr(k) }

func (p *Ctx) Json(m T) {
	if p.IsErr() {
		return
	}
	if b, err := json.Marshal(m); err != nil {
		p.Error(http.StatusInternalServerError, "InternalServerError")
	} else {
		p.Resp.Header().Set("Content-Type", "application/json; charset=utf-8")
		p.Resp.Write(b)
	}
}

func (p *Ctx) NotFound() {
	p.Error(http.StatusNotFound, "")
}
