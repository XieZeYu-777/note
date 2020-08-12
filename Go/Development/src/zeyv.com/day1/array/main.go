package main

import "fmt"

func main() {
	// 数组学习
	arr := [5]int{1, 2, 3, 4, 5}
	// for i, v := range arr {
	// 	fmt.Println(v, i)
	// }
	// fmt.Println("fmt", arr, arr[len(arr)-1])
	x := f1(&arr)
	fmt.Println(x)
}

func f(a [5]int) {
	fmt.Println(a, "func f")
}

func f1(a *[5]int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return sum
}
