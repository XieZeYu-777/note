package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController)GET()  {
	this.TplName = "login.html"
}


func (this *LoginController)Post()  {
	//this.TplName = "login.html"
	//o := orm.NewOrm()
	//uname := this.Input().Get("name") // 获取input name
	//o.Insert() // todo 数据库-插入数据
	//o.Delete() // todo 数据库-删除数据
}

