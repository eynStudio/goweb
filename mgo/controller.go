package mgo

import (
	. "github.com/eynstudio/gobreak"
	. "github.com/eynstudio/goweb"
)

type MgoController struct {
	Ctx    Context `di`
	Req    Req     `di`
	Params Values  `di`
}

func (p *MgoController) Id() GUID  { return GUID(p.Params.GetVal("id")) }
func (p *MgoController) Id1() GUID { return GUID(p.Params.GetVal("id1")) }
func (p *MgoController) Id2() GUID { return GUID(p.Params.GetVal("id2")) }

func (p *MgoController) HasId() bool  { return p.Params.GetVal("id") != "" }
func (p *MgoController) HasId1() bool { return p.Params.GetVal("id1") != "" }
func (p *MgoController) HasId2() bool { return p.Params.GetVal("id2") != "" }

func (p *MgoController) UserId() GUID {
	uid:= p.Ctx.Get(TYPE_GOWEB_USER_ID).Interface().(GOWEB_USER_ID)
	return GUID(uid)
}
