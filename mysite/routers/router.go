package routers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"github.com/zykzhang/site-demo/mysite/cache/myredis"
	"github.com/zykzhang/site-demo/mysite/controllers"
	// "golang.org/x/net/context"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
	beego.Any("/redis", func(ctx *context.Context) {
		glog.Infoln("redis....")
		fmt.Println("redis...")
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
