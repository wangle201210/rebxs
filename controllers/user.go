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
	qs := o.QueryTable("User")
	all, e := qs.All(&modUserList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{200,"获取成功！",modUserList}
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
