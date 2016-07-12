package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController1 struct {
	BaseController
}

type MainController2 struct {
	beego.Controller
}

func (this *MainController1) Post() {
	path := this.Ctx.Input.Param(":path")
	ext := this.Ctx.Input.Param(":ext")
	if path == "" {
		fmt.Println(1)
	}
	if ext == "" {
		fmt.Println(2)
	}
	this.Ctx.WriteString(path + " " + ext)
}

func (c *MainController1) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = id + ".tpl"
}

func (this *MainController2) Post() {
	splat := this.Ctx.Input.Param(":splat")
	if splat == "" {
		fmt.Println(1)
	}

	this.Ctx.WriteString(splat)
}
