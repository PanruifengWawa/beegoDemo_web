package models

import (
	//	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego/orm"
)

type UserToValidate struct {
	Id   int    `valid:"Min(1)"`
	Name string `valid:"Required"`
	Age  int    `valid:"Range(10,100)"`
}

func init() {
	orm.RegisterModel(new(UserToValidate))
}
