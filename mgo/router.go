package mgo

import (
	"fmt"
	"strings"

	. "github.com/eynstudio/goweb"
	"gopkg.in/mgo.v2/bson"
)

type USER_ID interface{}

var NilUserId *USER_ID

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

	ctrlInfo := r.FindController(route, params)
	ctx.Map(ctrlInfo)

	ctrl := &MgoController{}
	ctx.Apply(ctrl)
	ctx.Map(ctrl)

	return true
}

func MgoAuthHandler(ctx Context, ctrlInfo *CtrlInfo, req Req) bool {
	if !ctrlInfo.Auth {
		return true
	}

	jbreak := req.Header.Get("Authorization")
	if jbreak != "" {
		token := strings.Split(jbreak, " ")[1]
		if bson.IsObjectIdHex(token) {
			ctx.MapAs(token,NilUserId)
			return true
		}
	}
	ctx.Forbidden()
	return false
}
