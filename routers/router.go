package routers

import (
	"firstBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/index", &controllers.MainController{})
}
