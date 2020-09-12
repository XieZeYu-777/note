package controllers

import (
	"BeeBlog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	//this.Data["CategoryList"] =
	op := this.Input().Get("op")
	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
 		err := models.AddCategory(name)
		if err == nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	}
	fmt.Println()
	this.TplName = "category.html"
	var err error
	this.Data["CategoryList"], err = models.GetCategoryAll()
	if err != nil {
		beego.Error(err)
	}
}