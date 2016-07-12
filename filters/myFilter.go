package filters

import (
	"github.com/astaxie/beego/context"
)

func MyFilter(ctx *context.Context) {
	var uid int
	ctx.Input.Bind(&uid, "uid")
	if uid == 0 {
		ctx.Redirect(302, "/?token=testToken")
	}

}
