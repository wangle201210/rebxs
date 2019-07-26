package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/wangle201210/rebxs/models"
	"strconv"
)

type RewardController struct {
	BaseController
}

type RewardOne struct {
	Id       	int64		`json:"id"`
	User_id  	int64  		`json:"user_id"`
	Participant string		`json:"participant"`
	Type 		string    	`json:"type"`
	Time     	string    	`json:"time"`
}

var modReward models.Reward
var modRewardList []models.Reward

func (this *RewardController) All() {
	o := orm.NewOrm()
	qs := o.QueryTable("Reward")
	all, e := qs.All(&modRewardList)
	if e != nil {
		beego.Info(e)
	} else {
		beego.Info(all)
		resp = Response{200,"获取成功！",modRewardList}
	}
	this.Data["json"] = resp
	this.ServeJSON()
	return
}


func (this *RewardController) Save() {
	var Reward RewardOne
	var err error
	o := orm.NewOrm()
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &Reward); err == nil {
		user := models.User{Id: Reward.User_id}
		_ = o.Read(&user)
		in := models.Reward{User: &user, Participant: Reward.Participant,Type:Reward.Type,Time: Reward.Time}
		err := in.Insert()
		if err == nil {
			resp = Response{addSuccess.code,addSuccess.text,Reward}
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

func (this *RewardController) ShowList() {
	timeNow := this.GetString("time")
	user_id, _ := this.GetInt64("user_id")
	list := RewardList(user_id, timeNow)
	this.Data["json"] = list
	this.ServeJSON()
}

func RewardList(user_id int64,timeNow string) []orm.Params  {
	var maps []orm.Params
	o := orm.NewOrm()
	qs := o.QueryTable("reward")
	if timeNow != "" {
		qs = qs.Filter("time", timeNow)
	}
	if user_id != 0 {
		qs = qs.Filter("user_id",user_id)
	}
	fmt.Println("info: ",user_id,timeNow)
	_, _ = qs.Values(&maps,"user__name","participant","time","type","user__id")
	return maps
}

func RewardScore(user_id int64,timeNow string) (j []map[string]interface{}) {
	list := RewardList(user_id, timeNow)
	i := make(map[string]map[string]interface{})
	for _,row := range list{
		d := make(map[string]interface{})
		d["成员"] = row["User__Name"]
		d["项目"] = "悬赏"
		d["成绩"] =  "--"
		d["时间"] =  row["Time"]
		var score = 0
		switch row["Type"] {
		case "s":
			score = 1
		case "s+":
			score = 2
		case "ss":
			score = 4
		case "ss+":
			score = 4

		default:
			score = 0
		}
		if i[row["User__Name"].(string)]["分数"] != nil {
			s,_ := strconv.Atoi(i[row["User__Name"].(string)]["分数"].(string))
			s = s + score
			d["分数"] = strconv.Itoa(s)
		} else  {
			d["分数"] = strconv.Itoa(score)
		}
		i[row["User__Name"].(string)] = d // 用于计算
	}
	//转为slice
	for _,row := range i {
		j = append(j, row)
	}
	return j
}

