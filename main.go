package main

import (
	"beegoWeb/controllers"
	_ "beegoWeb/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	beego.ErrorController(&controllers.ErrorController{}) //my error defined
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Run()
}
