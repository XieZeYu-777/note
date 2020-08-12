package main

import "fmt"

func split(sum int) (x int, y int) {
	x = sum * 4
	y = sum - 1
	return
}

func main() {
	// 函数多种方式如下 第一个括号是传参 第二个是返回值（类型 支持闭包
	// c, d := func(a int, b int) (c int, d int) {
	// 	c = a + b
	// 	d = 4
	// 	return
	// }(1, 2)
	// fmt.Println(c, d)
	// c := func(a int, b int) (c int) {
	// 	c = a + b
	// 	return
	// }(1, 2)
	// fmt.Println(c)

	// c, d := func(a int, b int) (int, int) {
	// 	return a, b
	// }(1, 2)
	// fmt.Println(c, d)

	// ↓
	defer fmt.Println(split(10))

	defer fmt.Println("222")
	fmt.Println("111")

	// ↑ 执行顺序是从下到上的 输出 111 ，222， 40 9

	// 指针
	i1 := 1
	pint := &i1
	fmt.Println("i1", i1)
	data := *pint
	fmt.Println("data", data)
}
