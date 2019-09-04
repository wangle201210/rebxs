package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       	int64		`json:"id" orm:"column(id);auto;"`
	Name     	string    	`json:"name" orm:"column(name);size(100)"`
	Record      []*Record	`orm:"reverse(many)"`
	CreatedAt   time.Time	`json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time	`json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (reg *User) FindOrCreat (){
	read := reg.Read("Name")
	if read != nil {
		if reg.Name != "" {
			e := reg.Insert()
			if e != nil {
				fmt.Println(e)
			}
		}
	}
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