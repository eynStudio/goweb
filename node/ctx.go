package node

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
)

type CtxErr struct {
	Code int
	Msg  string
}

type Ctx struct {
	di.Container
	Req
	*Resp
	//	handlers []Handler
	Scope  M
	Err    *CtxErr
	afters []Handler
	urlParts
}

func (p *Ctx) Error(s int, d string) *Ctx {
	p.WriteHeader(s)
	p.Err = &CtxErr{Code: s, Msg: d}
	return p
}

func (p *Ctx) Set(k string, v T)   { p.Scope[k] = v }
func (p *Ctx) IsErr() bool         { return p.Err != nil }
func (p *Ctx) Get(k string) string { return p.Scope.GetStr(k) }
