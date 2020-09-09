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
	// 如果已经登录过的话 退出然后清空cookie
	isExits := this.Input().Get("exits") == "true"
	fmt.Println(isExits, "-isExit 16")
	if isExits {
		fmt.Println("go-isExit 18")
		// 清空cookie
		this.Ctx.SetCookie("unames", "", -1, "/")
		this.Ctx.SetCookie("pass", "", -1, "/")
		// 重定向
		fmt.Println("go-isExit 21")
		this.Redirect("/", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	unames := this.Input().Get("unames")
	pass := this.Input().Get("pass")
	authLogin := this.Input().Get("authLogin") == "on"
	// 判断是否和conf里配置一样 一样的话存cookie 重定向到首页
	if beego.AppConfig.String("unames") == unames &&
		beego.AppConfig.String("pass") == pass {
		maxAge := 0
		if authLogin {
			maxAge = 1 << 31 - 1
		}
		// 设置cookie
		this.Ctx.SetCookie("unames", unames, maxAge, "/")
		this.Ctx.SetCookie("pass", pass, maxAge, "/")
	}
	this.Redirect("/", 301)
	return
	//this.TplName = "login.html"
	//o := orm.NewOrm()
	//unames := this.Input().Get("name") // 获取input name
	//o.Insert() // todo 数据库-插入数据
	//o.Delete() // todo 数据库-删除数据
}

func ExitWith (ctx *context.Context) bool {
	// 获取cookie
	ck,err := ctx.Request.Cookie("unames")
	if err != nil {
		fmt.Println(err, "-Err Msg 58")
		return false
	}
	unames := ck.Value
	ck,err = ctx.Request.Cookie("pass")
	if err != nil {
		fmt.Println(err, "-Err Msg 64")
		return false
	}
	pass := ck.Value
	fmt.Println(beego.AppConfig.String("unames") == unames &&
		beego.AppConfig.String("pass") == pass, "EditWith 69")
	return beego.AppConfig.String("unames") == unames &&
		beego.AppConfig.String("pass") == pass
}

