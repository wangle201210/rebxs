package models

import (
	"github.com/astaxie/beego/orm"
)

type Record struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	User  		*User  		`orm:"rel(fk)"`
	Project     *Project    `orm:"rel(fk)"`
	Result 		int64    	`json:"result" orm:"column(result);size(100)"`
	Time     	string    	`json:"time" orm:"column(time);size(100)"`
}
// 多字段唯一键
func (u *Record) TableUnique() [][]string {
	return [][]string{
		[]string{"user_id", "project_id","time"},
	}
}
// Record database CRUD methods include Insert, Read, Update and Delete
func (reg *Record) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *Record) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Record) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Record) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(Record))
}