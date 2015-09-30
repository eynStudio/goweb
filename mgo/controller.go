package mgo

import (
	. "github.com/eynstudio/goweb"
	"github.com/eynstudio/gobreak/di"
	"gopkg.in/mgo.v2/bson"
)

type MgoController struct {
	Ctx    Context `di`
	Req    Req     `di`
	Params Values  `di`
}

func (p *MgoController) Id() bson.ObjectId  { return bson.ObjectIdHex(p.Params.GetVal("id")) }
func (p *MgoController) Id1() bson.ObjectId { return bson.ObjectIdHex(p.Params.GetVal("id1")) }
func (p *MgoController) Id2() bson.ObjectId { return bson.ObjectIdHex(p.Params.GetVal("id2")) }

func (p *MgoController) HasId() bool  { return bson.IsObjectIdHex(p.Params.GetVal("id")) }
func (p *MgoController) HasId1() bool { return bson.IsObjectIdHex(p.Params.GetVal("id1")) }
func (p *MgoController) HasId2() bool { return bson.IsObjectIdHex(p.Params.GetVal("id2")) }

func (p *MgoController) UserId() bson.ObjectId{
	uid:=p.Ctx.Get(di.InterfaceOf(NilUserId)).Interface().(string)
	return bson.ObjectIdHex(uid)	
}