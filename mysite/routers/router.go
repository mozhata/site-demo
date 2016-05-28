package routers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/golang/glog"
	"github.com/zykzhang/site-demo/mysite/cache/myredis"
	"github.com/zykzhang/site-demo/mysite/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
	beego.Any("/redis", func(ctx *context.Context) {
		glog.Infoln("line-20--redis....")
		fmt.Println("line21--redis...")
		tryRedis()
		ctx.Output.Body([]byte("reidis"))
	})
}
func tryRedis() {
	client := myredis.NewClient()
	pong, err := client.Ping().Result()
	glog.Infoln(pong, err)
	fmt.Println(pong, err)
}
