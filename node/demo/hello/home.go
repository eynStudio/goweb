package main

import (
	"log"
	"reflect"

	"github.com/eynstudio/goweb/node"
)

type Home struct {
	*node.Node
}

func (p *Home) Router(c *node.Ctx) {
	log.Printf("router: %#v", p)
	p.RunInterceptors(c)
	//	cur := c.CurPart()

	//	p.tryParseParam(c)
	p.RouteSubNodes(c)

	if !c.Handled {
		log.Println("node: ", p.Path, " not handled")
	}
	log.Println(c.Scope)
	//try handle here...
	p.Handler(c)
}

func (p *Home) Handler(c *node.Ctx) {
	log.Println("Handler at Home。。。。。。。。。。。。。")
	if c.Handled {
		log.Println("node: at path:", p.Path, ",ctx had handled")
	}
	log.Println("node:"+p.Path, " handler at home")
	v := reflect.TypeOf(p)
	cc := v.NumMethod()
	for i := 0; i < cc; i++ {
		m := v.Method(i)
		log.Println(m.Name)
	}
}

func (p *Home) Get(c *node.Ctx) {
	log.Println("Get...")
}
