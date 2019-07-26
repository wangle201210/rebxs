package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/wangle201210/rebxs/models"
	"strconv"
)

type UserController struct {
	BaseController
}

var modUser models.User
var modUserList []models.User

func (this *UserController) All() {
	o := orm.NewOrm()
	var lists []orm.Params

	_, err := o.Raw("SELECT * from `user` order by convert(name using gbk) asc").Values(&lists)


	if err != nil {
		beego.Info(err)
	} else {
		//beego.Info(lists)
		resp = Response{200,"获取成功！",lists}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *UserController) Add() {
	var User models.User
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &User); err == nil {
		err := User.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,User}
		} else {
			resp = Response{addError.code,addError.text,""}
		}
	} else {
		resp = Response{addError.code,addError.text,""}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *UserController) Delete() {
	id := this.Ctx.Input.Param(":id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		resp = Response{paraError.code,paraError.text,""}
	} else {
		modUser.Id = intId
		e := modUser.Delete()
		if e != nil {
			resp = Response{deleteError.code,deleteError.text,""}
		} else {
			resp = Response{deleteSuccess.code,deleteSuccess.text,""}
		}

	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}
