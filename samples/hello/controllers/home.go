package controllers

import (
	. "github.com/eynstudio/goweb"
)

type Home struct {
	Controller
}

func (this *Home) Get() Result {
	data := struct{ Name string }{"EYN"}
	return JsonResult{data}
}

func (this *Home) GetXyz() Result {
	data := struct{ Name string }{"XYZ"}
	return JsonResult{data}
}
