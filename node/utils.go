package node

import (
	"net/http"
	"strings"
)

type Handler func(c *Ctx)

type Interceptor struct {
	Before Handler
	After  Handler
}

type Req struct {
	*http.Request
}

func (p *Req) Url() string {
	return p.URL.Path
}

func (p *Req) Method() string {
	return strings.ToLower(p.Request.Method)
}

func (p *Req) IsJsonContent() bool {
	return strings.Contains(p.Header.Get("Content-Type"), "application/json")
}

func (p *Req) IsAcceptJson() bool {
	return strings.Contains(p.Header.Get("Accept"), "application/json")
}

type Resp struct {
	http.ResponseWriter
}

type urlPart struct {
	path string
}

type urlParts struct {
	curIdx int
	parts  []*urlPart
}

func newUrlParts(path string) *urlParts {
	m := &urlParts{}
	m.parseParts(path)
	return m
}

func (p *urlParts) parseParts(path string) {
	parts := strings.Split(path, "/")
	for _, it := range parts {
		p.parts = append(p.parts, &urlPart{it})
	}
}
func (p *urlParts) moveNextPart()     { p.curIdx += 1 }
func (p *urlParts) hasNextPart() bool { return p.curIdx < len(p.parts)-1 }
func (p *urlParts) CurPart() *urlPart { return p.parts[p.curIdx] }
func (p *urlParts) NextPart() *urlPart {
	if p.hasNextPart() {
		return p.parts[p.curIdx+1]
	}
	return nil
}
