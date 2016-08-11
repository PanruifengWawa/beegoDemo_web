package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["beegoWeb/controllers:AController"] = append(beego.GlobalControllerRouter["beegoWeb/controllers:AController"],
		beego.ControllerComments{
			"Haha",
			`/a/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["beegoWeb/controllers:AController"] = append(beego.GlobalControllerRouter["beegoWeb/controllers:AController"],
		beego.ControllerComments{
			"Hehe",
			`/a/:id:int`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["beegoWeb/controllers:ValidateController"] = append(beego.GlobalControllerRouter["beegoWeb/controllers:ValidateController"],
		beego.ControllerComments{
			"Hehe",
			`/`,
			[]string{"post"},
			nil})

}
