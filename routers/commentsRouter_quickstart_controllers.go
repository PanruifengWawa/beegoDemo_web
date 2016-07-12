package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["quickstart/controllers:AController"] = append(beego.GlobalControllerRouter["quickstart/controllers:AController"],
		beego.ControllerComments{
			"Haha",
			`/a/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["quickstart/controllers:AController"] = append(beego.GlobalControllerRouter["quickstart/controllers:AController"],
		beego.ControllerComments{
			"Hehe",
			`/a/:id:int`,
			[]string{"post"},
			nil})

}
