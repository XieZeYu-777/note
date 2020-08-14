package main

import "fmt"

func main()  {
	// todo 指针不能用来运算
	var count = 20
	var countPoint *int // 指针类型的变量
	countPoint = &count // &去指针地址
	fmt.Printf("指针地址,%x", countPoint)
	fmt.Printf("指针地址,%d", *countPoint) // 去指针地址值
	// 指针数组
	var a,b = 1,2
	pointArr := [...] *int{&a,&b}
	fmt.Println(pointArr, "指针数组")
	// 数组指针
	arrPointArr := &[...]int{1,2,3}
	fmt.Println(arrPointArr, "数组指针")

}
