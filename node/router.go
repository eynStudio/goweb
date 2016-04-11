package node

import (
	"log"
)

type Router struct {
}

func (p *Router) Route(n INode, c *Ctx) {
	n.RunInterceptors(c)
	p.RouteSubNodes(n, c)

	if !c.Handled {
		n.Handler(c)
	}
}

func (p *Router) RouteSubNodes(n INode, c *Ctx) {
	if c.hasNextPart() {
		for _, it := range n.GetNodes() {
			log.Println("next node path ", c.NextPart().path, it)
			if it.CanRouter(c.NextPart().path) {
				log.Printf("can route %#v", it)
				c.moveNextPart()
				p.Route(it, c)
				break
			} else {
				log.Printf("can not route %#v", it)
			}
		}
	}
}
