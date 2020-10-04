package main

import "fmt"

func main() {

	// 指针

	// &是取地址 *是取地址值

	var n = 1
	var a = &n

	fmt.Println(a, *a, test(1, &n))

}

func test(a int, b *int) int {
	return a + *b
}
