package controllers

import (
	"github.com/astaxie/beego"
)

type FlashController struct {
	beego.Controller
}

func (this *FlashController) Get() {
	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["json"] = n
	} else if n, ok := flash.Data["error"]; ok {
		this.Data["json"] = n
	} else {
		this.Data["json"] = "no flash"
	}
	this.ServeJSON()
}

func (this *FlashController) Post() {
	flash := beego.NewFlash()
	id, _ := this.GetInt("id")
	if id == 0 {
		flash.Error("id is null")
		flash.Store(&this.Controller)
		this.Redirect("/flashTest", 302)
		return
	}
	flash.Notice("success")
	flash.Store(&this.Controller)
	this.Redirect("/flashTest", 302)

}
