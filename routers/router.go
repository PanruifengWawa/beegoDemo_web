package routers

import (
	//	"fmt"
	"beegoWeb/controllers"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/context"

	"beegoWeb/filters"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Get("/api1", func(ctx *context.Context) {
		ctx.Output.Body([]byte("body"))
	})
	//basic router

	beego.Router("/api2/?:id", &controllers.MainController{})
	//id is not necessary, if there is no id, then it will be ""

	beego.Router("/api3/:id", &controllers.MainController{})
	//id is necessary, if there is no id, then 404

	beego.Router("/api4/:id([0-9]+)", &controllers.MainController{})
	beego.Router("/api5/:id:int", &controllers.MainController{})
	//id is necessary and id is a number

	beego.Router("/api6/:id:string", &controllers.MainController{})
	//id is necessary and id is a string

	beego.Router("/api7/*.*", &controllers.MainController1{})
	//param : :path and :ext

	beego.Router("/api8/*", &controllers.MainController2{})
	//param : :splat

	beego.Router("/api9/view_:id([0-9]+).html", &controllers.MainController1{})
	beego.Router("/api10/view_:id:int.html", &controllers.MainController1{})
	beego.Router("/api11/view_:id:string.html", &controllers.MainController1{})

	beego.Router("/test", &controllers.UdController{}, "get,delete:Test1;post:Test2;put:Test3")

	beego.AutoRouter(&controllers.UdController{})

	beego.Include(&controllers.AController{})

	ns := beego.NewNamespace("/apiA",
		beego.NSCond(func(ctx *context.Context) bool {
			if ctx.Input.Domain() == "localhost" {
				return true
			}
			return false
		}),
		beego.NSRouter("/default", &controllers.MainController{}),

		beego.NSNamespace("/my",
			beego.NSInclude(&controllers.AController{}),
		),
	)
	beego.AddNamespace(ns)

	beego.InsertFilter("/apiA/my/*", beego.BeforeRouter, filters.MyFilter)
	//filter /api/my/* with function 'MyFilter' in package filters

	beego.Router("/flashTest", &controllers.FlashController{}, "post:Post;get:Get")
	beego.Router("/validate", &controllers.ValidateController{}, "post:Hehe;get:Haha")

	//	beego.Router("/errors", &controllers.ErrorController{})
}
