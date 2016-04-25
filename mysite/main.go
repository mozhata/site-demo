package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zykzhang/site-demo/mysite/routers"
)

func main() {
	beego.SetStaticPath("/sta", "temp")
	beego.Router("/hello", &MainController{})
	beego.Run()
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}
