package node

import (
	"log"
)

type INode interface {
	Node(path string) *Node
	Router(c *Ctx) (found bool)
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

type ParamNode struct{ *Node }

func NewParamNode(path string) *ParamNode { return &ParamNode{NewNode(path)} }

func (p *Node) Node(path string) *Node {
	n := NewNode(path)
	p.Nodes = append(p.Nodes, n)
	return n
}

func (p *Node) AddNode(n INode) *Node {
	p.Nodes = append(p.Nodes, n)
	return p
}

func (p *Node) Router(c *Ctx) (found bool) {
	p.RunInterceptors(c)

	log.Println("Node.Router")
	return false
}

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
