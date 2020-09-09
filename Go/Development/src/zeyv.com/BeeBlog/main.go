package main

import (
	"BeeBlog/models"
	_ "BeeBlog/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init()  {
	// 数据库初始化
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	// todo 数据库名称 是否一直重新建表 是否打印相关信息 开发默认都是true
	orm.RunSyncdb("default", false, true)
	beego.Run()
}

