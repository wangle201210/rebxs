package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/wangle201210/rebxs/models"
)

type ProjectController struct {
	BaseController
}

var modProject models.Project
var modProjectList []models.Project

func (this *ProjectController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("Project")
	all, e := qs.All(&modProjectList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{200,"获取成功！",modProjectList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *ProjectController) Add() {
	var Project models.Project
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Project); err == nil {
		err := Project.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Project}
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

