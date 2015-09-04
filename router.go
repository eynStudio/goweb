package goweb

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	Fixed = iota
	Parameter
	Any
)

var MyRouter Router

type Route struct {
	Source       string
	SourcePath   string
	SourceSearch string
	Segments     []string
	Prefix       string
	Regexp       string
	Params       []string
	ParamNames   []string
}

type Router struct {
	Routes []Route
	Ctrls  []interface{}
}

func NewRoute(path string) (route *Route) {
	route = &Route{Source: path}
	route.parse()
	return
}

func (this *Route) parse() {
	placeholder := `([:*])([\w\[\]]+)|\{([\w\[\]]+)(?:\:((?:[^{}\\]+|\\.|\{(?:[^{}\\]+|\\.)*\})+))?\}`
	//	searchPlaceholder := `([:]?)([\w\[\]-]+)|\{([\w\[\]-]+)(?:\:((?:[^{}\\]+|\\.|\{(?:[^{}\\]+|\\.)*\})+))?\}`
	reg := regexp.MustCompile(placeholder)

	pattern := this.Source
	quoteRegExp := func(seg string, pattern string, optional bool) (result string) {
		surroundPattern := []string{"(", ")"}
		result = strings.Replace(seg, `[\\\[\]\^$*+?.()|{}]`, "\\$&", -1)
		if len(pattern) == 0 {
			return
		}
		if optional {
			surroundPattern = []string{"(", ")?"}
		}
		return result + surroundPattern[0] + pattern + surroundPattern[1]
	}
	last := 0
	all := reg.FindAllStringSubmatchIndex(pattern, -1)
	compiled := "^"
	fmt.Printf("%q\n", reg.FindAllStringSubmatch(pattern, -1))
	fmt.Println(all)
	for _, r := range all {
		segment := pattern[last:r[0]]
		this.Segments = append(this.Segments, segment)

		compiled = compiled + quoteRegExp(segment, "[^/]*", false)

		if r[4] > -1 {
			this.ParamNames = append(this.ParamNames, pattern[r[4]:r[5]])
		} else {
			this.ParamNames = append(this.ParamNames, pattern[r[6]:r[7]])
		}

		last = r[1]
	}
	segment := pattern[last:]

	compiled = compiled + quoteRegExp(segment, "", false) + "$"

	this.Regexp = compiled
	this.Segments = append(this.Segments, segment)
	this.SourcePath = pattern
	this.Prefix = this.Segments[0]
	fmt.Printf("%#v\n", this)
}

func (this *Route) Exec(path string) {
	reg := regexp.MustCompile(this.Regexp)
	all := reg.FindAllStringSubmatch(path, -1)
	fmt.Println(all)
}

func (this *Router) Route(path string) {
	r := NewRoute(path)
	r.Exec("/api/abc-xyz/123")

	this.Routes = append(this.Routes, *r)
}

func (this *Router) Register(controller *Controller) {
	this.Ctrls = append(this.Ctrls, controller)
}
