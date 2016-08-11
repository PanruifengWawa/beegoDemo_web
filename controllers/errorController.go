package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.TplName = "404.tpl"
	c.Data["content"] = "not found"
}

func (this *ErrorController) ErrorDb() {
	this.TplName = "Db.tpl"
	this.Data["content"] = "db error"

}

func (this *ErrorController) ErrorHh() {
	this.TplName = "Db.tpl"
	this.Data["content"] = "haha"

}
