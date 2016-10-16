package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"

	c.Data["Email"] = "zyk7676@gmail.com"

	c.TplName = "index.tpl"
}
