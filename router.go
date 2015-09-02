package goweb

import (
	"fmt"
	"regexp"
)

const (
	Fixed = iota
	Parameter
	Any
)

var MyRouter Router

type RouteSegment struct {
	Type int
	Name string
}

type RouteSegments []RouteSegment

type Matcher struct {
	Source       string
	SourcePath   string
	SourceSearch string
	Segments     []string
	Prefix       string
	Regexp       string
	Params       []string
	ParamNames   []string
}
type Route struct {
	Url      string
	Segments []RouteSegments
}

type Router struct {
	Routes []Route
	Ctrls  []interface{}
}

func (this *Router) Route(path string) {
	r := &Route{Url: path}
	r.ParseUrl()
	this.Routes = append(this.Routes, *r)
}

func (this *Router) Register(controller *Controller) {
	this.Ctrls = append(this.Ctrls, controller)
}

func (this *Route) ParseUrl() {
	fmt.Println(this.Url)
	UrlMatcher(this.Url)
}

func UrlMatcher(pattern string) {
	var matcher Matcher
	matcher.Source = pattern

	placeholder := `([:*])([\w\[\]]+)|\{([\w\[\]]+)(?:\:((?:[^{}\\]+|\\.|\{(?:[^{}\\]+|\\.)*\})+))?\}`
	//	searchPlaceholder := `([:]?)([\w\[\]-]+)|\{([\w\[\]-]+)(?:\:((?:[^{}\\]+|\\.|\{(?:[^{}\\]+|\\.)*\})+))?\}`
	reg := regexp.MustCompile(placeholder)

	last := 0
	all := reg.FindAllStringSubmatchIndex(pattern, -1)
	fmt.Printf("%q\n", reg.FindAllStringSubmatch(pattern, -1))
	fmt.Println(all)
	for _, r := range all {
		segment := pattern[last:r[0]]
		matcher.Segments = append(matcher.Segments, segment)
		if r[4] > -1 {
			matcher.ParamNames = append(matcher.ParamNames, pattern[r[4]:r[5]])
		} else {
			matcher.ParamNames = append(matcher.ParamNames, pattern[r[6]:r[7]])
		}

		last = r[1]
	}
	segment := pattern[last:]
	matcher.Segments = append(matcher.Segments, segment)
	matcher.SourcePath=pattern
	matcher.Prefix=matcher.Segments[0]
	fmt.Printf("%#v\n", matcher)
}
