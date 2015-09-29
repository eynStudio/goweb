package mgo

import (
	"fmt"
	. "github.com/eynstudio/goweb"
	//	"gopkg.in/mgo.v2/bson"
)

func MgoRouterHandler(ctx Context, r Router, req Req) bool {
	url := req.Url()
	fmt.Println(url)

	route, params := r.FindRoute(url)
	if route == nil {
		ctx.Json("route not found")
		return false
	}
	ctx.Map(params)
	ctx.Map(route)

	ctrl := &MgoController{}
	ctx.Apply(ctrl)
	ctx.Map(ctrl)

	return true
}
