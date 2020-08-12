package main

import (
	"fmt"
)

func returnVal(a, b *int, c *[]int) {
	*a = 10
	*b = 20
	*c = []int{2, 3, 4, 5}
}

//	自定义函数类型
type FuncTest func(int, int) int

type FunT func(int int) int

func main() {
	a := 100
	b := 200
	/*
	* TODO 值类型直接修改无效 因为值类型存储的是值不是指针 引用类型存储的是指针地址 要指针修改
	* 值类型有int bool string strut 数组 float
	 */
	switch {
	case a == 100:
		fmt.Println(100)
		fallthrough
	case a == 200:
		fmt.Println(23)
	default:
		fmt.Println("33")
	}

	c := make([]int, 1)

	returnVal(&a, &b, &c)
	fmt.Println(a, c[0])

	fmt.Println(test01(3))

	// var d FuncTest

	// d = Add

	// res := d(1, 3)
	// fmt.Println(res)

	// var aaa FunT

	// aaa = Add1

	// res1 := aaa(333)

	// fmt.Println(res1, "函数类型返回")

	fff1 := Add222(1,2, Add)
	fmt.Println(fff1)

}

func Add222 (a, b int, iFff FuncTest) (vc int) {
	vc = Add(a,b)
	return 
}

func Add1(a int) int {
	return a
}
func Add(a, b int) int {
	return a + b
}

func test01(i int) int {
	fmt.Println(i, "i")
	if i == 1 {
		return 1
	}
	return i + test01(i-1)
}
