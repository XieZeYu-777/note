package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

type User struct {
	Name string
	Age int
}
func (c *HomeController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsBool"] = false
	var Nums = []int {1,2,3}
	c.Data["Nums"] = Nums
	c.Data["IsLogin"] = ExitWith(c.Ctx)
	c.TplName = "index.tpl"
	// 结构体
	u := User{Name:"Yo",Age:10}
	// 赋值
	c.Data["User"] = u
	// 模板变量定义
	c.Data["tplVar"] = "hi guys "
	c.Data["Html"] = "<div>html div </div>"
	c.Data["Pipe"] = "<div>html Pipe </div>"

}
