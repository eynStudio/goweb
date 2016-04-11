package node

type ParamNode struct {
	*Node
}

func NewParamNode(path string) *ParamNode { return &ParamNode{NewNode(path)} }

func (p *ParamNode) CanRouter(test string) bool { return true }

func (p *ParamNode) Handler(c *Ctx) {
	c.Scope[p.Path] = c.CurPart().path
}
