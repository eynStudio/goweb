package node

import (
	"regexp"
)

type RegexNode struct {
	*Node
	regex string
}

func NewRegexNode(path, regex string) *RegexNode { return &RegexNode{Node: NewNode(path), regex: regex} }

func (p *RegexNode) CanRouter(test string) bool {
	match, _ := regexp.MatchString(p.regex, test)
	return match
}

func (p *RegexNode) Handler(c *Ctx) {
	c.Scope[p.Path] = c.CurPart().path
}
