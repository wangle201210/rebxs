package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	"github.com/wangle201210/rebxs/helper"
	"github.com/wangle201210/rebxs/models"
	"strconv"

	//"github.com/wangle201210/rebxs/helper"
	"time"
)

type RecordController struct {
	BaseController
}

type RecordList struct {
	Id       	int64		`json:"id"`
	User_id  	int64  		`json:"user_id"`
	Project_id  int64  		`json:"project_id"`
	Result 		int64    	`json:"result"`
	Time     	string    	`json:"time"`
}

type DataListRecord struct {
	uname	string
	pname 	string
	score 	int64
	result 	int64
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
	ob := []RecordList{}
	o := orm.NewOrm()
	e := json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	if e !=  nil {
		fmt.Println(e)
	}
	//fmt.Println(ob)
	//this.ServeJSON()
	//return
	one := ob[0]
	project := models.Project{Id: one.Project_id}
	err := o.Read(&project)
	if err != nil {
		fmt.Println("项目呢？？")
	}
	timeNow := getTime()
	for _, row := range ob {
		user := models.User{Id: row.User_id}
		err = o.Read(&user)
		in := models.Record{User: &user, Project: &project,Result:row.Result,Time: timeNow}
		_,err:=o.InsertOrUpdate(&in,"time","user_id","project_id")
		if err == nil {
			fmt.Println("插入成功")
		}
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

func  (this *RecordController) GetList()  {
	time := this.GetString("time")
	fmt.Println("time",time)
	o := orm.NewOrm()
	var lists []orm.Params
	num, err := o.Raw("SELECT u.name AS '成员',p.name as '项目',r.result as '成绩',s.score as '分数',r.time as '时间' FROM user u LEFT JOIN record r ON r.user_id = u.id LEFT JOIN project p ON p.id = r.project_id LEFT JOIN score s ON s.project_id = p.id and s.min <= r.result and r.result <= s.max WHERE r.id > 0  and r.time = ? ", time).Values(&lists)
	if err == nil && num > 0 {
	}
	var data []map[string]interface{}
	for _,row := range lists{
		//map[string]interface{}
		d := make(map[string]interface{})
		for k, v := range  row{
			d[k] = v
		}
		data = append(data,d)
	}
	res := make(map[string]interface{})
	title := helper.GetValueByKey(data, "项目")
	title = helper.SliceRemoveDuplicate(title)//获取全部成员
	r := DataFormat(data)
	rs := helper.MapToSlice(r)
	res["data"] = rs
	res["org"] = data
	res["title"] = title
	//fmt.Println(data)
	//rs := helper.DataFormatBykey(data,"项目")
	//fmt.Println(r)
	//bytes, err := json.Marshal(r)
	//if err != nil {
	//	fmt.Println("json err:",err)
	//}
	this.Data["json"] = res
	this.ServeJSON()
}

func DataFormat(d []map[string]interface{}) (r map[string]map[string]interface{})  {
	r3 := helper.GetValueByKey(d, "成员")
	r4 := helper.SliceRemoveDuplicate(r3)//获取全部成员
	rs := make(map[string]map[string]interface{})
	for _,row := range r4{
		var total int
		d4p := make(map[string]interface{})
		d4p["成员"] = row
		for _,dp := range d{
			if dp["成员"] == row{
				i, _ := strconv.Atoi(dp["分数"].(string))
				d4p[dp["项目"].(string)] = i
				total += i
			}
		}
		d4p["总分"] = total
		rs[row.(string)] = d4p
	}
	r = rs
	return
}

func (this *RecordController) ShowOne() {
	all := []models.Record{}
	timeNow := this.GetString("time")
	project_id := this.GetString("project_id")
	o := orm.NewOrm()
	qs := o.QueryTable("record")
	_, _ = qs.Filter("project_id", project_id).Filter("time",timeNow).All(&all)
	fmt.Println("ShowOne",timeNow,project_id,all)
	this.Data["json"] = all
	this.ServeJSON()
}
