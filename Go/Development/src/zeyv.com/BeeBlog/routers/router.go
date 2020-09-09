package routers

import (
	"BeeBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &controllers.LoginController{})
}
