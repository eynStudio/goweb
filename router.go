package goweb

import (
	"fmt"
	//	"path"
	"reflect"
	"regexp"
	"strings"
)

const (
	Fixed = iota
	Parameter
	Any
)

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
	Ctrls  map[string]*ControllerInfo
}

func NewRouter() (router *Router) {
	return &Router{Ctrls: make(map[string]*ControllerInfo)}
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
		surroundPattern := []string{"(", ")?"}
		result = strings.Replace(seg, `[\\\[\]\^$*+?.()|{}]`, "\\$&", -1)
		if len(pattern) == 0 {
			return
		}
		//		if optional {
		//			surroundPattern = []string{"(", ")?"}
		//		}
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

		if r[4] > -1 {
			this.ParamNames = append(this.ParamNames, pattern[r[4]:r[5]])
		} else {
			this.ParamNames = append(this.ParamNames, pattern[r[6]:r[7]])
		}
		patt := "[^/]*"
		if r[8] > -1 {
			patt = pattern[r[8]:r[9]]
		}
		compiled += quoteRegExp(segment, patt, false)

		last = r[1]
	}
	segment := pattern[last:]

	compiled += quoteRegExp(segment, "", false) + "$"

	this.Regexp = compiled
	this.Segments = append(this.Segments, segment)
	this.SourcePath = pattern
	this.Prefix = this.Segments[0]
	fmt.Printf("%#v\n", this)
}

func (this *Route) Exec(path string) (bool, map[string]string) {
	reg := regexp.MustCompile(this.Regexp)
	if !reg.MatchString(path) {
		return false, nil
	}

	all := reg.FindStringSubmatch(path)[1:]
	params := make(map[string]string, len(all))

	for i, param := range all {
		params[this.ParamNames[i]] = param
	}
	return true, params
}

func (this *Router) Route(path string) {
	r := NewRoute(path)
	r.Exec("/api/abc-xyz/123")
	r.Exec("/api/abc-xyz/opq")

	this.Routes = append(this.Routes, *r)
}

func (this *Router) Register(controller interface{}) {
	c := reflect.TypeOf(controller)
	fmt.Println(c)
	name := strings.ToLower(c.Elem().Name())

	ctrl := NewControllerInfo(name, c.Elem())
	this.Ctrls[name] = ctrl

	for i, j := 0, c.NumMethod(); i < j; i++ {
		m := c.Method(i)
		ctrl.Methods[strings.ToLower(m.Name)] = ControllerMethod{m.Name, nil}
	}
}

func (this *Router) FindRoute(url string) (*Route, map[string]string) {
	for _, r := range this.Routes {
		match, params := r.Exec(url)
		if match {
			fmt.Printf("%#v\n", params)
			return &r, params
		}
	}
	return nil, nil
}
func (this *Router) FindController(route *Route, params map[string]string) *ControllerInfo {
	ctrl := params["ctrl"]
	ctrlInfo := this.Ctrls[ctrl]
	return ctrlInfo
}

func RouterFilter(ctx *HttpContext, fc []Filter) {
	url := ctx.Req.URL.Path
	fmt.Println(url)
	ctx.Route, ctx.Params = ctx.App.Router.FindRoute(url)
	if ctx.Route == nil {
		ctx.Result = &JsonResult{"route not found"}
	} else {
		fc[0](ctx, fc[1:])
	}
}

func ControllerFilter(ctx *HttpContext, fc []Filter) {
	ctrlInfo := ctx.App.Router.FindController(ctx.Route, ctx.Params)
	execController(ctrlInfo, ctx)
	fc[0](ctx, fc[1:])
}

func execController(ctrlInfo *ControllerInfo, ctx *HttpContext) {
	ctrl := reflect.New(ctrlInfo.Type)
	method := strings.ToLower(ctx.Req.Method)
	var act ControllerMethod
	if action, ok := ctx.Params["action"]; ok {
		if m, ok := ctrlInfo.Methods[method+action]; ok {
			act = m
		}
	} else if id, ok := ctx.Params["id"]; ok && id != "" {
		if m, ok := ctrlInfo.Methods[method+"id"]; ok {
			act = m
		}
	} else if m, ok := ctrlInfo.Methods[method]; ok {
		act = m
	}

	if act.Name != "" {
		fmt.Println(act.Name)
		ctx.Result = ctrl.MethodByName(act.Name).Call(nil)[0].Interface().(Result)
		fmt.Println(ctx.Result)
		fmt.Println(ctrl.Elem())
	} else {
		fmt.Println("No Action!")
	}
}
