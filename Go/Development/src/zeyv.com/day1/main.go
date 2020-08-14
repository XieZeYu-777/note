package main

// 导入包名
import (
	"fmt"
)

// 变量和常量

// 单独声明变量
// var str string // 声明一个变量必须要声明数据类型
// 声明变量go官方推荐驼峰式命名
//
// 批量声明
//var (
//	name string
//	age  int
//	isOk bool
//)


type FunMain func (int int) int



// 例子1
//func foo() (int, string) {
//	return 11, "22"
//}

// todo 声明变量必须使用 全局的没事函数里声明的变量的必须使用
// a := "1"
// b := 1
// a = b 变量类型不同不能赋值
func main() {
	//importStruct.TestStruct()
	// 匿名变量_下划线 用于处理函数返回多个 你只想用一个用来占位 不会占用内存 如 例子1
	//x, _ := foo()
	//_, y := foo()
	//fmt.Println("x的值是", x)
	//fmt.Println("y的值是", y)
	//// 函数里赋值
	//name = "zeyv"
	//age = 22
	//isOk = true
	//// 变量 类型推导 根据声明的变量值判断是什么数据类型
	//var s2 = "123"
	//// 简短变量声明:="hhh!"
	//s3 := "hhh!"                  // 是var s3 = "hhh"的简写 只能在函数里使用~
	//fmt.Println(s2, s3)           // 打印完后面加一个换行符
	//fmt.Printf("name:%s\n", name) // \n 换行 %s 相当于占位符使用这个name变量的值替换占位符字符串类型 %v 替换所有类型 int string 都可以 用法一样
	//// fmt.Printf("%#v", s3) =
	//// fmt 是引入的库
	//fmt.Println("hello world 你he好世界") // 打印完后面加一个换行符
	// s1 := "22" // 同一个作用域不能声明重复的变量

	// go语言中字符串双引号包裹 "时代大厦sdd"
	// go语言中字符单引号包裹 'h' '1' '是'


	var a FunMain
	fmt.Println("函数类型", a(1))



}
