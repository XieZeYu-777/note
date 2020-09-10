package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get()  {
	// 清空cookie
	//this.Ctx.SetCookie("uname", "", -1, "/")
	//this.Ctx.SetCookie("pass", "", -1, "/")
	fmt.Println("golang peace", this.Input().Get("exit") == "true")
	IsEdit := this.Input().Get("exit") == "true"
	if IsEdit {
		// 清空cookie
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pass", "", -1, "/")
		// 重定向
		this.Redirect("/",301)
		return
	}
	this.TplName = "login.html"
}


func (this *LoginController) Post() {
	uname := this.Input().Get("uname")
	pass := this.Input().Get("pass")
	authLogin := this.Input().Get("authLogin") == "on"
	fmt.Println(uname, pass, "NotConfig")
	fmt.Println(beego.AppConfig.String("uname"), beego.AppConfig.String("pass"), "YesConfig")
	// 判断是否和conf里配置一样 一样的话存cookie 重定向到首页
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pass") == pass {
		maxAge := 0
		if authLogin {
			maxAge = 1 << 31 - 1
		}
		// 设置cookie
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pass", pass, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
}

func IfLogin (c *context.Context) bool {
	// 获取cookie
	uname := c.GetCookie("uname")
	pass := c.GetCookie("pass")
	// 如果是nil返回false
	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pass") == pass
}


