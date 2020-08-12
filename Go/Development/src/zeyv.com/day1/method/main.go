package main

import "fmt"

type TZ int // 类型别名

type Num int // 类型别名

type A struct {
	Name string
	name string // 大小写区别 小写只能这个文件访问到，大写整个包都能访问到
}

type B struct {
	Name string
}

func main () {
	var tz TZ // 初始化
	// TODO Method Value
	tz.Print()
	// TODO Method Expression
	(*TZ).Print(&tz)
	a:= A{}
	a.Print(1,2)
	fmt.Println(a) // {"AA", "A123"}

	b:= B{}
	b.Print(3,4)
	fmt.Println(b) // {}

	var num Num // 初始化
	num.Sum100(100)
	fmt.Println(num, "Sum100")
}
func (num *Num) Sum100 (a int) {
	*num += Num(a)
}
=
func (tz TZ) Print () {
	fmt.Println("TZ")
}

func (a *A) Print (arg1 ,arg2 int) {
	a.Name = "AA" // 指针修改
	a.name = "A123" // 这样是改变不了结构体的
	fmt.Println("A", arg1, arg2)
}

func (b B) Print (arg1 ,arg2 int) {
	b.Name = "BB" // 这样是改变不了结构体的
	fmt.Println("B", arg1, arg2)
}
