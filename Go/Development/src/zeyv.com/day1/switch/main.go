package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

func main() {
	var g int
	go func(j int) {
		a := 0
		for i := 0; i < j; i++ {
			fmt.Println(i, j)
			a += i
			g = a
		}
	}(1000)

	//  计算函数执行时间
	start := time.Now()
	test()
	end := time.Now()
	dele := end.Sub(start)
	fmt.Printf("dele=%d\n", dele)
	// map类型

	var mapData map[int]string

	mapData = make(map[int]string)

	mapData[0] = "zeyv"

	fmt.Printf("map[0] = %s\n", mapData[0])

	// var a int = 1
	// switch a {
	// case 1:
	// 	fmt.Println("a1")
	// case 2:
	// 	fmt.Println("a2")
	// case 3:
	// 	fmt.Println("a3")
	// default:
	// 	fmt.Println("a4")
	// }

	// var b int = 1
	// // 如果在执行完每个分支的代码还想继续执行的话用fallthrough
	// switch {
	// case b >= 1:
	// 	fmt.Println("b1")
	// 	fallthrough
	// case b >= 2:
	// 	fmt.Println("b2")
	// 	fallthrough
	// case b >= 3:
	// 	fmt.Println("b3")
	// 	fallthrough
	// default:
	// 	fmt.Println("b4")
	// }
	// // 简短变量 switch
	// switch c := 1; c {
	// case 1:
	// 	fmt.Println("C1")
	// case 2:
	// 	fmt.Println("C2")
	// default:
	// 	fmt.Println("default")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("THIS is %d\n", i)
	// }

	// for i := 0; i < 5; i++ {
	// 	var v int
	// 	fmt.Printf("%d 1111111111111111111", v)
	// 	v = 5
	// }

	// range 类似其他语言的forEach 获取他的value 和 index 如下

	// var str = "日本语"

	// for index, val := range str {
	// 	fmt.Println(val, index, len(str))
	// }

	// str := "g"

	// for i1 := 0; i1 < 25; i1++ {
	// 	str += "g"
	// 	println(str)
	// }
	// 用goto形式输出0-14 不用for

	// 	i1 := 0

	// START:
	// 	fmt.Println(i1)
	// 	i1++
	// 	if i1 < 15 {
	// 		goto START
	// 	}

	// 	fmt.Printf("sum a * b * c = %d\n", sum(1, 2, 3))
	// a := 10
	// b := 20

	// //swap01(a, b)     //值传递（传值）
	// swap02(&a, &b)                       //地址传递（传引用）
	// fmt.Printf("a = %d, b = %d\n", a, b) // 20 10
	deFers()

	where := func() {
		_, file, line, _ := runtime.Caller(1)
		fmt.Println(file, line)
	}
	where()
	ppp := Hz(".jpg")

	fmt.Println(ppp("fil1111e"))
}

func test() string {
	return "ss"
}

// 工厂函数

func Hz(str string) func(a string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, str) {
			name += str
		}
		return name
	}
}

func swap01(a, b int) {
	a, b = b, a
	fmt.Printf("swap01 a = %d, b = %d\n", a, b)
}

func swap02(x, y *int) {
	// *取地址值 &地址传递
	*x, *y = *y, *x
	fmt.Println(*x, *y, x, y)
}

// defer 类似栈 后进先出

func deFers() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("defer i = %d\n", i)
	}
}

// func sum(a, b, c int) int {
// 	return a * b * c
// }
