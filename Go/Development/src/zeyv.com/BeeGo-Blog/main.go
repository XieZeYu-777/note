package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get () {
	this.Ctx.WriteString("hello wrold ")
}

func main()  {
	// 注册路由
	beego.Router("/", &HomeController{})
	// 启动服务
	beego.Run()
}
