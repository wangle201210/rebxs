package models

import (
	"github.com/astaxie/beego/orm"
)

type Project struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	Name     	string    	`json:"name" orm:"column(name);size(100)"`
	Record		[]*Record	`orm:"reverse(many)"`
}

// Project database CRUD methods include Insert, Read, Update and Delete
func (reg *Project) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *Project) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Project) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Project) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(Project))
}