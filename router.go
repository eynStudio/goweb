package goweb

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/eynstudio/gobreak/di"
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
	DefaultCtrl  string
}

type Router interface {
	RegisterAuth(controller interface{})
	RegisterAnon(controller interface{})
	Route(path string)
	RouteCtrl(path, ctrlName string)
	FindRoute(url string) (*Route, Values)
	FindController(route *Route, params Values) *CtrlInfo
}
type router struct {
	Routes []Route
	Ctrls  map[string]*CtrlInfo
}

func NewRouter() Router {
	return &router{Ctrls: make(map[string]*CtrlInfo)}
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
	if this.DefaultCtrl != "" {
		vals.Set("ctrl", this.DefaultCtrl)
	}
	return true, vals
}

func (this *router) Route(path string) {
	r := NewRoute(path)
	this.Routes = append(this.Routes, *r)
}
func (this *router) RouteCtrl(path, ctrlName string) {
	r := NewRoute(path)
	r.DefaultCtrl = ctrlName
	this.Routes = append(this.Routes, *r)
}
func (this *router) RegisterAuth(controller interface{}) {
	this.register(controller, true)
}
func (this *router) RegisterAnon(controller interface{}) {
	this.register(controller, false)
}

func (this *router) register(controller interface{}, needAuth bool) {
	c := reflect.TypeOf(controller)
	name := strings.ToLower(c.Elem().Name())

	ctrl := NewCtrlInfo(name, c.Elem(), needAuth)
	this.Ctrls[name] = ctrl

	for i, j := 0, c.NumMethod(); i < j; i++ {
		m := c.Method(i)
		ctrl.Methods[strings.ToLower(m.Name)] = CtrlAction{m.Name, di.GetMethodArgs(m.Type)}
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

func (this *router) FindController(route *Route, vals Values) *CtrlInfo {
	ctrl, _ := vals.Get("ctrl")
	if ctrl == "" {
		ctrl = "default"
	}
	ctrlInfo := this.Ctrls[ctrl]
	return ctrlInfo
}

func RouterHandler(ctx Context, r Router, req Req) bool {
	url := req.Url()
	fmt.Println(req.Method() + " : " + url)

	route, params := r.FindRoute(url)
	if route == nil {
		ctx.Json("route not found")
		return false
	}
	ctx.Map(params)
	ctx.Map(route)

	ctrlInfo := r.FindController(route, params)
	ctx.Map(ctrlInfo)

	baseCtrl := &Controller{}
	ctx.Apply(baseCtrl)
	ctx.Map(baseCtrl)

	return true
}

func CtrlHandler(ctx Context, req Req, ctrlInfo *CtrlInfo, params Values) bool {
	method := req.Method()
	jBreakMethod := req.Header.Get("jBreak-Method")
	if jBreakMethod != "" {
		jBreakMethod = strings.ToLower(jBreakMethod)
	}

	var actions []string
	act2, ok2 := params.Get("act2")
	act1, ok1 := params.Get("act1")
	_, ok := params.Get("id")

	appendActions := func(act string) {
		if jBreakMethod != "" {
			actions = append(actions, method+act+jBreakMethod)
		}
		actions = append(actions, method+act)
	}

	if ok2 {
		appendActions(strings.ToLower(act2))
	}
	if ok2 && ok1 {
		appendActions(strings.ToLower(act1) + strings.ToLower(act2))
	}
	if ok1 {
		appendActions(strings.ToLower(act1))
	}
	if ok {
		appendActions("id")
	}
	appendActions("")

	for _, act := range actions {
		if m, ok := ctrlInfo.Methods[act]; ok {
			ctx.Map(m)
			return true
		}
	}
	fmt.Println("No Action!")
	return false
}
func BindingHandler(ctx Context, req Req, act CtrlAction) bool {
	if req.IsJsonContent() && req.Body != nil && len(act.Args) > 0 && act.Args[0].Kind() != reflect.Interface {
		defer req.Body.Close()

		data := reflect.New(act.Args[0])
		err := json.NewDecoder(req.Body).Decode(data.Interface())
		if err != nil {
			fmt.Println(err)
		} else {
			ctx.Map(data.Elem().Interface())
		}
	}
	return true
}
func ActionHandler(ctx Context, ctrlInfo *CtrlInfo, act CtrlAction) bool {
	ctrl := reflect.New(ctrlInfo.Type)
	ctx.Apply(ctrl.Interface())

	if m := ctrl.MethodByName("Before"); m.IsValid() {
		ctx.Exec(m, nil)
	}

	ctx.Exec(ctrl.MethodByName(act.Name), act.Args)

	if m := ctrl.MethodByName("After"); m.IsValid() {
		ctx.Exec(m, nil)
	}

	return true
}
