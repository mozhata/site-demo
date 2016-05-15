package routers

import (
	"github.com/astaxie/beego"
	"github.com/zykzhang/site-demo/mysite/controllers"
	// "golang.org/x/net/context"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
}
