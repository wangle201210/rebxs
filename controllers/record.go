package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/wangle201210/rebxs/models"
	"time"
)

type RecordController struct {
	BaseController
}

type RecordList struct {
	Id       	int64		`json:"id"`
	User_id  	[]int64  	`json:"user_id"`
	Project_id  int64  		`json:"project_id"`
	Result 		int64    	`json:"result"`
	Time     	string    	`json:"time"`
}

var modRecord models.Record
var modRecordList []models.Record

func (this *RecordController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("Record")
	all, e := qs.All(&modRecordList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{200,"获取成功！",modRecordList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *RecordController) Add() {
	var Record models.Record
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Record); err == nil {
		err := Record.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Record}
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



func (this *RecordController) Save() {
	//s := getTime()
	//fmt.Println(s)
	//return
	var ob = RecordList{}
	//var record []*models.Record
	o := orm.NewOrm()
	qs := o.QueryTable("record")
	i, _ := qs.PrepareInsert()
	e := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if e !=  nil {
		fmt.Println(e)
	}
	project := models.Project{Id: ob.Project_id}

	err := o.Read(&project)
	if err != nil {
		fmt.Println("项目呢？？")
	}
	for _, user_id := range ob.User_id {
		user := models.User{Id: user_id}
		err = o.Read(&user)
		fmt.Println(user,project)
		in := models.Record{User: &user, Project: &project,Result: ob.Result,Time: getTime()}
		id, err := i.Insert(&in)
		if err == nil {
			fmt.Println("插入成功",id)
		}
	}

	ec := i.Close() // 别忘记关闭 statement

	if ec != nil {
		fmt.Println("关闭失败",ec)
	}

	this.Data["json"] = ob
	this.ServeJSON()
}

func getTime() (s string)  {
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	now := time.Now().Unix()
	var jg int64
	jg = 604800
	startTime := appConf.String("startTime")
	p, _ := time.Parse("2006-01-02 15:04:05", startTime)
	unix := p.Unix()
	cz := now - unix
	z := cz / jg + 1
	s = fmt.Sprintf("第%d周", z)
	return
}