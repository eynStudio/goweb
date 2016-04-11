package main

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/goweb/node"
)

type Api struct {
	*node.Node
}

func NewApi() *Api {
	h := &Api{node.NewNode("api")}
	h.NewParamNode("id")
	return h
}

func (p *Api) Handler(c *node.Ctx) {
	handled := true
	switch c.Method {
	case "GET":
		p.Get(c)
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *Api) Get(c *node.Ctx) {
	c.Json(M{"haah": "aa"})
}
