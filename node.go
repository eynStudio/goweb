package goweb

type Node struct {
	Path  string
	Nodes []*Node
}

func NewNode(path string) *Node {
	return &Node{
		Path:  path,
		Nodes: []*Node{},
	}
}
func (p *Node) Sub(path string) *Node {
	n := NewNode(path)
	p.Nodes = append(p.Nodes, n)
	return n
}
