package main

import (
	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/goweb/node"
	"log"
)

type Home struct {
	*node.Node
}

func NewHome() *Home {
	h := &Home{node.NewNode("", false)}
	h.NewParamNode("id", false)
	return h
}

func (p *Home) Handler(c *node.Ctx) {
	handled := true
	switch c.Method {
	case "GET":
		p.get(c)
	case "POST":
		p.post(c)
	default:
		handled = false
	}
	c.Handled = handled
}

func (p *Home) get(c *node.Ctx) {
	c.Json(M{"get": "aa"})
}

func (p *Home) post(c *node.Ctx) {
	var h H
	c.JsonBody(&h)
	log.Println(h)
	c.Json(M{"post": "aa"})
}

type H struct {
	Id int
}
