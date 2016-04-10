package node

import (
	"log"
)

type INode interface {
	NewNode(path string) INode
	AddNode(n INode) INode
	CanRouter(test string) bool
	Router(c *Ctx)
}

type Node struct {
	Path         string
	Interceptors []*Interceptor
	Nodes        []INode
}

func NewNode(path string) *Node {
	return &Node{
		Path:         path,
		Interceptors: []*Interceptor{},
		Nodes:        []INode{},
	}
}

func (p *Node) NewNode(path string) INode {
	n := NewNode(path)
	p.Nodes = append(p.Nodes, n)
	return n
}

func (p *Node) AddNode(n INode) INode {
	p.Nodes = append(p.Nodes, n)
	return p
}

func (p *Node) Router(c *Ctx) {
	log.Printf("router: %#v", p)
	p.RunInterceptors(c)
	//	cur := c.CurPart()

	p.tryParseParam(c)
	p.RouteSubNodes(c)

	if !c.Handled {
		log.Println("node: ", p.Path, " not handled")
	}
	log.Println(c.Scope)
	//try handle here...
	p.Handler(c)
}

func (p *Node) Handler(c *Ctx) {
	if c.Handled {
		log.Println("node: at path:", p.Path, ",ctx had handled")
	}
	log.Println("node:"+p.Path, " handler here")
}
func (p *Node) RouteSubNodes(c *Ctx) {
	if c.hasNextPart() {
		for _, it := range p.Nodes {
			log.Println("next node path ", c.NextPart().path, it)
			if it.CanRouter(c.NextPart().path) {
				log.Printf("can route %#v", it)
				//				log.Println()
				c.moveNextPart()
				it.Router(c)
				break
			} else {
				log.Printf("can not route %#v", it)
			}
		}
	} else {
		log.Println(p.Path, " no next path")
	}
}

func (p *Node) tryParseParam(c *Ctx) {
	if p.isParamNote() {
		c.Scope[p.getParamName()] = c.CurPart().path
	}
}

func (p *Node) CanRouter(test string) bool { return p.isParamNote() || p.Path == test }
func (p *Node) isParamNote() bool {
	return len(p.Path) > 0 && '{' == p.Path[0] && '}' == p.Path[len(p.Path)-1]
}
func (p *Node) getParamName() string { return p.Path[1 : len(p.Path)-1] }
func (p *Node) isRegexNode() bool    { return '(' == p.Path[0] && ')' == p.Path[len(p.Path)-1] }

func (p *Node) Interceptor(m *Interceptor) *Node {
	p.Interceptors = append(p.Interceptors, m)
	return p
}

func (p *Node) RunInterceptors(c *Ctx) *Node {
	if c.IsErr() {
		return p
	}

	for _, i := range p.Interceptors {
		if nil != i.After {
			c.afters = append(c.afters, i.After)
		}

		if nil != i.Before {
			i.Before(c)
			if c.IsErr() {
				break
			}
		}
	}

	return p
}

type ParamNode struct {
	*Node
}

func NewParamNode(path string) *ParamNode { return &ParamNode{NewNode(path)} }
