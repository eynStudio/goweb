package goweb

import (
	"fmt"
	"github.com/eynstudio/gobreak/di"
	"net/http"
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

type Router interface {
	Register(controller interface{})
	Route(path string)
	FindRoute(url string) (*Route, Values)
	FindController(route *Route, params Values) *ControllerInfo
}
type router struct {
	Routes []Route
	Ctrls  map[string]*ControllerInfo
}

func NewRouter() Router {
	return &router{Ctrls: make(map[string]*ControllerInfo)}
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

func (this *Route) Exec(path string) (bool, Values) {
	reg := regexp.MustCompile(this.Regexp)
	if !reg.MatchString(path) {
		return false, nil
	}

	all := reg.FindStringSubmatch(path)[1:]
	vals := make(Values, len(all))

	for i, param := range all {
		vals.Set(this.ParamNames[i], param)
	}
	return true, vals
}

func (this *router) Route(path string) {
	r := NewRoute(path)
	r.Exec("/api/abc-xyz/123")
	r.Exec("/api/abc-xyz/opq")

	this.Routes = append(this.Routes, *r)
}

func (this *router) Register(controller interface{}) {
	c := reflect.TypeOf(controller)
	fmt.Println(c)
	name := strings.ToLower(c.Elem().Name())

	ctrl := NewControllerInfo(name, c.Elem())
	this.Ctrls[name] = ctrl

	for i, j := 0, c.NumMethod(); i < j; i++ {
		m := c.Method(i)
		ctrl.Methods[strings.ToLower(m.Name)] = ControllerMethod{m.Name, di.GetMethodArgs(m.Type)}
	}
}

func (this *router) FindRoute(url string) (*Route, Values) {
	for _, r := range this.Routes {
		match, vals := r.Exec(url)
		if match {
			fmt.Printf("%#v\n", vals)
			return &r, vals
		}
	}
	return nil, nil
}

func (this *router) FindController(route *Route, vals Values) *ControllerInfo {
	ctrl, _ := vals.Get("ctrl")
	ctrlInfo := this.Ctrls[ctrl]
	return ctrlInfo
}

func RouterHandler(ctx Context, r Router, req *http.Request) bool {
	url := req.URL.Path
	fmt.Println(url)

	route, params := r.FindRoute(url)
	if route == nil {
		ctx.Json("route not found")
		return false
	}
	ctx.Map(params)
	ctx.Map(route)
	return true
}

func CtrlHandler(ctx Context, req *http.Request, r Router, route *Route, params Values) bool {
	ctrlInfo := r.FindController(route, params)
	ctx.Map(ctrlInfo)

	method := strings.ToLower(req.Method)
	var act ControllerMethod
	if action, ok := params.Get("action"); ok {
		if m, ok := ctrlInfo.Methods[method+action]; ok {
			act = m
		}
	} else if id, ok := params.Get("id"); ok && id != "" {
		if m, ok := ctrlInfo.Methods[method+"id"]; ok {
			act = m
		}
	} else if m, ok := ctrlInfo.Methods[method]; ok {
		act = m
	}

	if act.Name == "" {
		fmt.Println("No Action!")
		return false
	}

	ctx.Map(act)
	fmt.Println(act.Name)
	return true
}
func BindingHandler() bool {

	return true
}
func ActionHandler(ctx Context, ctrlInfo *ControllerInfo, act ControllerMethod) bool {
	ctrl := reflect.New(ctrlInfo.Type)
	ctx.Exec(ctrl.MethodByName(act.Name), act.Args)
	fmt.Println(ctx.(*context).Result)
	return true
}
