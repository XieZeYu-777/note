package main

import "fmt"

func main() {
	a := 1
	//  if判断流
	if a == 1 {
		fmt.Println("a =  1")
	} else if a == 2 {
		fmt.Println("a = 2")
	}
	// go可以在if里声明变量 如下
	if a := 1; a == 1 {
		fmt.Println("a == 1")
	}

	// for 循环
	// for a := 0; a < 10; a++ {
	// 	fmt.Println(a)
	// }
	// go 没有while 用for解决

	b := 0
	for b != 10 {
		b++
		fmt.Println(b)
	}

	var obj map[string]string // 定义map

	obj = make(map[string]string)

	obj["name"] = "zeyv"
	obj["age"] = "23"

	fmt.Println(obj)
	for k, v := range obj {
		//循环体
		fmt.Println(k, v)
	}

	// i := 1
	// // Here:
	// if i == 1 {
	// 	i = 2
	// 	// goto Here
	// }

	// fmt.Println(i, "here")

	// list := []int{1, 2, 3, 4}

	// i := 2
	// switch i {
	// case 1, 2:
	// 	fmt.Println("123")
	// 	fallthrough // 相当于break 这样下面就可以不用写break类似
	// case 3, 4:
	// 	fmt.Println("345")
	// }

	// for val, index := range list {
	// 	fmt.Println(val, index)
	// }
}
