package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/glog"
	_ "github.com/zykzhang/site-demo/server/routers"
)

func main() {
	beego.SetStaticPath("/sta", "temp")
	beego.Router("/hello", &MainController{})

	beego.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})
	beego.Any("/redis", func(ctx *context.Context) {
		glog.Infoln("line-20--redis....")
		tryRedis2()
		// tryRedis()
		ctx.Output.Body([]byte("reidis"))
	})

	beego.Run()
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world~~ ")
}
func tryRedis2() {
	// conn := myredis.NewConn()
	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	v, err := conn.Do("SET", "name", "red")
	if err != nil {
		glog.Error(err)
	}
	glog.Infoln("v: ", v)
	v, err = redis.String(conn.Do("GET", "name"))
	if err != nil {
		glog.Error(err)
	}
	glog.Info("value by get: ", v)
}
