package main

import (
	"log"

	"github.com/eynstudio/goweb/node"
)

type Home struct {
	*node.Node
}

func (p *Home) Handler(c *node.Ctx) {
	log.Println("Handler at Home。。。。。。。。。。。。。")
	if c.Handled {
		log.Println("node: at path:", p.Path, ",ctx had handled")
	}
	log.Println("node:"+p.Path, " handler at home")
}
