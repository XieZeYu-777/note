package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func main() {
	p := add(2)

	fmt.Printf("add = %d\n", p(3))

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
		fmt.Printf("g=%d\n", g)
	}(1000)

	ppp := Hzstrings(".jpg")

	fmt.Println(ppp("file")) // file.jpg

	where := func() {
		_, flie, line, _ := runtime.Caller(1)
		log.Print("%s:%d", flie, line)
	}
	where() // 2020/04/28 22:02:48 %s:%dF:/Go/Development/src/zeyv.com/day1/returnfunc/main.go33
	where() // 2020/04/28 22:02:48 %s:%dF:/Go/Development/src/zeyv.com/day1/returnfunc/main.go33
	where() // 2020/04/28 22:02:48 %s:%dF:/Go/Development/src/zeyv.com/day1/returnfunc/main.go33

}

func add(a int) func(b int) int {
	return func(b int) int {
		fmt.Printf("b = %d\n", b)
		fmt.Printf("a = %d\n", a)
		return a + b
	}
}

// 工厂函数
func Hzstrings(su string) func(a string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, su) {
			name += su
		}
		return name
	}
}

func Hz(str string) func(b string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, str) {
			name += str
		}
		return name
	}
}
