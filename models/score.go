package models

import (
	"github.com/astaxie/beego/orm"
)

type Score struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	Project  	*Project  	`orm:"rel(fk)"`
	Min     	int64    	`json:"min" orm:"column(min);size(100)"`
	Max     	int64    	`json:"max" orm:"column(max);size(100)"`
	Score     	int64    	`json:"score" orm:"column(score);size(100)"`
}

// Score database CRUD methods include Insert, Read, Update and Delete
func (reg *Score) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *Score) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Score) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *Score) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(Score))
}