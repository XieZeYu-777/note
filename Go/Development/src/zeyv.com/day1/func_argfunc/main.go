package main

import "fmt"

func main() {
	callback(3, add)
	// 闭包
	// func() {
	// 	sum := 0
	// 	for i := 0; i < 10; i++ {
	// 		sum += i
	// 	}
	// 	fmt.Printf("sum=%d\n", sum)
	// }()

	s := func(a, b int) int { return a + b }
	fmt.Println(s(19, 1))

	n := 1
	test(&n, 1)

}

func test(a *int, b int) int {
	fmt.Println(a, b, *a)
	*a = 2
	fmt.Println(a, b, *a, "修改")
	return *a + b
}

func add(a, b int) {
	fmt.Println(a, b, a+b)
	//  a 3
	//  b 1
}

func callback(y int, f func(int, int)) {
	f(y, 1)
}
