package routers

import (
	"flag"

	"github.com/zykzhang/site-demo/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	flag.Parse()
	beego.Router("/", &controllers.MainController{})
	// beego.Any("/foo", func(ctx *context.Context) {
	// 	ctx.Output.Body([]byte("bar"))
	// })
	// beego.Any("/redis", func(ctx *context.Context) {
	// 	glog.Infoln("line-20--redis....")
	// 	tryRedis2()
	// 	// tryRedis()
	// 	ctx.Output.Body([]byte("reidis"))
	// })
}
