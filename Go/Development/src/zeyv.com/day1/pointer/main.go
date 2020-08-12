package main

import (
	"fmt"
	"unicode/utf8"
)

type TH int
type STR string

func main() {
	var a, b TH = 1, 2
	var e, d STR = "HELLO '\n'", "WORLD"
	c := a + b
	g := e + d
	fmt.Println(c, g)
	// 指针pointer &*
	fmt.Println("sss")

	s := "hel" + "lo,"
	s += "world"
	fmt.Println(s)

	str := "Sdsdsdf fdfdf"

	i := len([]byte(str))
	println(i)

	fmt.Printf("Length: %d, Runes:%d\n", i, utf8.RuneCount([]byte(str)))
	// 阻塞等待用户输入
	// var a int
	// fmt.Scanf("%d\n", &a)
	// var a int
	// fmt.Sprintf("谢泽雨%d\n", a)
	// // time库
	// var times = time.Now()
	// // fmt.Println(time.Sleep(time.Second)) // 每隔一秒执行
	// fmt.Println(times)
}
