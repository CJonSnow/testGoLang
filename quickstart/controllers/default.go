package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https//:XXX.com1"
	c.Data["Email"] = "astaxie@gmail.com2"
	c.TplName = "index.tpl"
}
