package routers

import (
	"github.com/astaxie/beego"
	"github.com/zykzhang/site-demo/mysite/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
