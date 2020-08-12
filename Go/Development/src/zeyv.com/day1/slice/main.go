package main

import (
	"fmt"
)

func main() {
	// s := []int{1, 2, 3, 4}
	a := 1
	n := 2
	add(&a, &n)
	// fmt.Printf("%T", s)
	// fmt.Printf("%v", s[0:2])
	// 二维数组 len 长度 cap容量
	var arr [2][2]int
	arr = [2][2]int{
		{1, 2},
		{3, 4},
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[j][i])
		}
	}
}

func add(a, b *int) {
	fmt.Println(a, *b)
	*b = 44
	fmt.Println(a, *b)
}
