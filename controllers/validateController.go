package controllers

import (
	"fmt"
	//	"encoding/json"
	//	"fmt"
	"beegoWeb/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type ValidateController struct {
	BaseController
}

func (this *ValidateController) Haha() {
	this.Abort("Db")
}

// @router / [post]
func (this *ValidateController) Hehe() {
	beego.Debug("hhaha")

	u := models.UserToValidate{}
	if err := this.ParseForm(&u); err == nil {
		valid := validation.Validation{}

		//		valid.Required(u.Name, "name")
		//		valid.Range(u.Age, 10, 30, "age")
		//		valid.Min(u.Id, 1, "id")
		//		if valid.HasErrors() {
		//			for _, err := range valid.Errors {
		//				fmt.Println(err.Key, err.Message)
		//			}
		//			this.Data["json"] = valid.Errors
		//		}

		//		v := valid.Range(u.Age, 10, 20, "age")
		//		if !v.Ok {
		//			fmt.Println(v.Error)
		//			this.Data["json"] = v.Error
		//		}

		ok, err := valid.Valid(&u)
		if !ok && err == nil {
			fmt.Println(valid.Errors)
			this.Data["json"] = valid.Errors
		}
	}

	this.ServeJSON()

}
