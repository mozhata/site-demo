package routers

import (
	"github.com/zykzhang/site-demo/mysite/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
