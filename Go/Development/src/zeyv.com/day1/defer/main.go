package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
	for i := 10; i >= 1; i-- {
		fmt.Printf("i= %d\n", i)
	}
	// entering b
	// in b
	// entering a
	// in a
	// leaving a
	// leaving b
}
