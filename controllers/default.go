package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://www.baidu.com/?tn=80035161_1_dg"
	c.Data["Email"] = "1342308281@qq.com"
	c.TplName = "index.tpl"
}
