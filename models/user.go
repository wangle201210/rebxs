package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	Name     	string    	`json:"name" orm:"column(name);size(100)"`
	Record      []*Record	`orm:"reverse(many)"`
}

func Find(m User) (r User) {
	read := m.Read("Name", "Password")
	if read != nil {
		beego.Debug(read)
	}
	r = m
	return
}
// User database CRUD methods include Insert, Read, Update and Delete
func (reg *User) Insert() error {
	if _, err := orm.NewOrm().Insert(reg); err != nil {
		return err
	}
	return nil
}

func (reg *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(reg, fields...); err != nil {
		return err
	}
	return nil
}

func (reg *User) Delete() error {
	if _, err := orm.NewOrm().Delete(reg); err != nil {
		return err
	}
	return nil
}


func init() {
	orm.RegisterModel(new(User))
}