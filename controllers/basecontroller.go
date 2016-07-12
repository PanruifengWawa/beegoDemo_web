package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	token := this.GetString("token")
	if token != "testToken" {
		this.Data["json"] = "auth error"
		this.ServeJSON()
		this.StopRun()
	}

}
