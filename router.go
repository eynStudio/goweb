package goweb

import (
	"fmt"
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
}
