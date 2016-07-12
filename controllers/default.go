package controllers

import (
	"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	//c.Ctx.WriteString("hahah")
	c.Render()
}

func (this *MainController) Post() {
	s := this.Ctx.Input.Param(":id")
	if s == "" {
		fmt.Println(1)
	}
	this.Ctx.WriteString(s)

}
