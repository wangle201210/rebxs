package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Reward struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	User  		*User  		`orm:"rel(fk)"`
	Type     	string    	`json:"type" orm:"column(type);size(100)"`
	Participant string    	`json:"participant" orm:"column(participant);size(100)"`
	Time     	string    	`json:"time" orm:"column(time);size(100)"`
	CreatedAt   time.Time	`json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time	`json:"updated_at" orm:"auto_now;type(datetime)"`
}
// Reward database CRUD methods include Insert, Read, Update and Delete
func (reg *Reward) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *Reward) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Reward) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Reward) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(Reward))
}