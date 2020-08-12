package main

import (
	"fmt"
	"os"
	"runtime"
)

// 常量 常量通常定义在全局
const pi = 233 // 不能被改变

// 批量声明常量

const (
	s1 = 1
	s2 = 2
)

// 如果批量声明常量时 某一行没有赋值默认和上一行一致

const (
	n1 = 100
	n2
	n3
)

// iota 常量计数器 iota当const出现de时候初始化为0 const中没新增一个将使iota加1
const (
	a1 = iota // 0
	a2        // = iota // 1
	a3        // = iota // 2
)

// 面试题

const (
	b1 = iota // 0
	b2 = 100  // 100  如果批量声明常量时 某一行没有赋值默认和上一行一致
	b3 = iota // 2 const中没新增一个将使iota加1
	b4 = iota // 3
)

const (
	c1, c2 = iota + 1, iota + 2 // c1 = 1, c2 = 2
	c3, c4 = iota + 1, iota + 2 // c3 = 2. c4 = 3
)

// << 向左平移
// %s 编译字符串
// %d 编译十进制
// 把十进制转换为二进制%b
// 把十进制转换为二进制%o
// 把十进制转换为二进制%x
// 查看变量类型%T
// 明确声明int8类型|int16 如 int8(变量名)，int16(变量名)
var (
	GOPATH = os.Getenv("GOPATH") // os.Getenv检测环境变量 没有返回空
)

const (
	z1 = iota // 0
	z2        // 100
	z3        // 2
)

func main() {
	var goods string = runtime.GOOS // 检测操作系统
	fmt.Println(goods)
	// 浮点数
	f1 := 1.2344 // go语言中默认浮点数是flota64位的 如果想转成32位的要用到float32(变量)
	fmt.Printf("%T\n", f1)
	fmt.Println(GOPATH)
	fmt.Println("z1", z1)
	fmt.Println("z2", z2)
	fmt.Println("z3", z3)
	// fmt.Println("a1", a1)
	// fmt.Println("a2", a2)
	// fmt.Println("a3", a3)

}
