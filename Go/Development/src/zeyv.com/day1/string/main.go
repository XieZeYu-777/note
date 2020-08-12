// package main

// import (
// 	"fmt"
// )

// func main() {
// 	a := 1

// 	fmt.Println(a)
// }

// package main

// var a = "G" // 全局变量

// func main() {
// 	n()
// 	m()
// 	n()
// }

// // 全局变量和局部变量
// func n() { print(a) }

// func m() {
// 	a = "O" // 局部变量
// 	print(a)
// }

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var a string

func main() {
	// str := " he,l,lo woRld,"
	// fmt.Println(strings.HasSuffix(str, "he"), strings.HasPrefix(str, "he"))
	// // strings.HasPrefix(str, "he") // 前缀是否还是str
	// // strings.HasSuffix(str, "he") // 后缀是否还是str
	// o := strings.Contains(str, "o")        //出现的位置 返回索引
	// z := strings.LastIndex(str, "l")       // 最后出现的位置 返回索引
	// s := strings.Replace(str, "l", "w", 1) // 替换Replace
	// // 统计字符串出现的次数
	// l := strings.Count(str, "l") // 3
	// fmt.Println(o, z, s, l)
	// // 重复字符串 返回新的字符串
	// p := strings.Repeat(str, 3)
	// fmt.Println(p)
	// e := strings.ToLower(str) // 转换成小写
	// r := strings.ToUpper(str) // 转换成大写
	// a = "G"
	// m := strings.Trim(str, "he")
	// fmt.Println("\n", m, "trim")
	// n := strings.Fields(str)
	// x := strings.Split(str, ",")

	// ji := strings.Join(n, ";")
	// fmt.Println("Join", ji)
	// fmt.Println("fields", n, x)

	// print(a, e, r)
	// f1()

	str := "你好,我是,谁"

	a := strings.Split(str, ",") // 返回一个slice
	fmt.Println(a)
	// 把slice数据拼接起来
	b := strings.Join(a, ";")
	fmt.Println(b) //你好;我是;谁

	// time库

	// t := time.Now()
	// UTC 国家时间标准
	// fmt.Printf("%02d.%02d.%04d", t.Day(), t.Month(), t.Year())
	// 指针

	// var i1 = 9
	// var i2 = &i1 // 取i1的地址
	// *i3 = 666    // i1地址值
	// fmt.Println(i2, i3)
	s := "sdsdsd"

	fmt.Printf("系统位数是： %d\n", strconv.IntSize)

	var p *string = &s     // 取地址 带类型的变量等要加*
	*p = "ciao"            // 改变指针内存地址值
	pp := *p               // 赋值
	fmt.Println(*p, s, pp) // ciao
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
