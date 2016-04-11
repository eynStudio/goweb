package main

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/goweb/node"
)

type Home struct {
	*node.Node
}

func NewHome() *Home {
	h := &Home{node.NewNode("")}
	h.NewParamNode("id")
	return h
}

func (p *Home) Handler(c *node.Ctx) {
	handled := true
	switch c.Method {
	case "GET":
		p.Get(c)
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *Home) Get(c *node.Ctx) {
	c.Json(M{"haah": "aa"})
}
