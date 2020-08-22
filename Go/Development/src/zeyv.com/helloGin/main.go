package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main () {
	// 实例
	router := gin.Default()
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()
	// 实例1 get，post，put，delete，patch,head,options
	//router.GET("/someGet",getting)

	// 获取路径中的参数
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Println(name, "name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		// todo message = name is / aciton
		fmt.Println(message, "msg")
		c.String(http.StatusOK, message)
	})
	// 获取get参数
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
		//lastname := c.DefaultQuery("lastname", "hi") // 是 c.Request.URL.Query().Get("lastname") 的简写
		c.String(http.StatusOK, "hello %s %s", firstname, lastname)
	})
	// 获取post参数
	router.POST("/from_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})
	// get+post混用
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		if id == "" {
			fmt.Println("error")
		}
		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id,page,name,message)
	})
	router.Run(":8080")
}

func getting(c *gin.Context)  {
	c.JSON(200, gin.H{
		"get": 200,
	})
}