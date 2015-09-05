package controllers

import (
	. "github.com/eynstudio/goweb"
)

type Home struct {
	Controller
}

func (this *Home) Get() Result {
	return TemplateResult{"index", "GoWeb"}
}

func (this *Home) GetXyz() Result {
	data := struct{ Name string }{"XYZ"}
	return JsonResult{data}
}
