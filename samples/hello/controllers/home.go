package controllers

import (
	. "github.com/eynstudio/goweb"
)

type Home struct {
}

func (this *Home) Get(ctx Context) {
	ctx.Tmpl("index", "GoWeb")
}

func (this *Home) GetXyz(ctx Context) {
	data := struct{ Name string }{"XYZ"}
	ctx.Json(data)
}
