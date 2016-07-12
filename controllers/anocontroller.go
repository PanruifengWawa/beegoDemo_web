package controllers

//import (
//	//	"encoding/json"
//	//	"fmt"
//	"quickstart/models"
//)

type AController struct {
	BaseController
}

// @router /a/:id [get]
func (this *AController) Haha() {
	//	this.Data["json"] = this.Ctx.Input.Param(":id")
	s := this.GetSession("mySession")
	if s == nil {
		this.SetSession("mySession", int(1))
		this.Data["json"] = 1
	} else {
		this.SetSession("mySession", s.(int)+1)
		this.Data["json"] = s.(int)
	}
	this.ServeJSON()

}

// @router /a/:id:int [post]
func (this *AController) Hehe() {
	this.Data["json"] = this.Ctx.Input.Param(":id")
	//	var s string = this.GetString("my")
	//	fmt.Println(s)

	//	u := models.User{}
	//	if err := this.ParseForm(&u); err == nil {
	//		fmt.Println(u)
	//		this.Data["json"] = u
	//	}

	//	var ob models.User
	//	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	//	this.Data["json"] = ob

	//	f, h, err := this.GetFile("myFile")
	//	defer f.Close()
	//	if err == nil {
	//		err2 := this.SaveToFile("myFile", "f:/"+h.Filename)
	//		if err2 == nil {
	//			this.Data["json"] = "success"
	//		} else {
	//			fmt.Println(err2)
	//			this.Data["json"] = "save error"
	//		}
	//	} else {
	//		this.Data["json"] = "file error"
	//	}

	//bind option with ?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
	//	var myInt int
	//	this.Ctx.Input.Bind(&myInt, "myInt")
	//	var myFloat float64
	//	this.Ctx.Input.Bind(&myFloat, "myFloat")
	//	var myBool bool
	//	this.Ctx.Input.Bind(&myBool, "myBool")
	//	myArr := make([]int, 0, 2)
	//	this.Ctx.Input.Bind(&myArr, "myArr")
	//	var user models.User
	//	this.Ctx.Input.Bind(&user, "user")
	//	this.Data["json"] = user

	this.ServeJSON()

}
