package main

import "fmt"

type Dog struct {
	Common // 组合 继承
	Name string
	Age int
	Sex string
}

type Common struct {
	eye string
}

func main () {
	// 初始化方式 一
	//var dog Dog
	//dog.Name = "点点"
	//dog.Age = 25
	//dog.Sex = "母"
	//fmt.Println(dog)
	// 初始化方式 二
	//Tstruct := Dog{Name:"帅仔",Age: 24, Sex: "男"}
	//fmt.Println(Tstruct)
	// 初始化方式 三
	dog := new(Dog) // 指向Dog类型的指针地址 他返回的是指针类型
	dog.Name = "Jock"
	dog.Age = 25
	dog.Sex = "母"
	dog.Common.eye = "my eyes"
	dog.Run()
	dog.Eat()
	dog.test()

	fmt.Println(dog)


	// todo 结构体的属性和函数 || 两种作用域
}
// todo method 函数
// todo Run 外部调用要开头大写才能调用成功 如果小写话只能在这个文件里访问 外部 别的调用不来的
func (d *Dog) Run () {
	fmt.Println(d.Name, d.Common.eye)
}

func (d *Dog) Eat () {
	fmt.Println(d.Age)
}

func (c * Common) test () {
	fmt.Println(c.eye)
}