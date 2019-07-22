package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/wangle201210/rebxs/models"
)

type ScoreController struct {
	BaseController
}

var modScore models.Score
var modScoreList []models.Score

func (this *ScoreController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("Score")
	all, e := qs.All(&modScoreList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{200,"获取成功！",modScoreList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *ScoreController) Add() {
	var Score models.Score
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Score); err == nil {
		err := Score.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Score}
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


