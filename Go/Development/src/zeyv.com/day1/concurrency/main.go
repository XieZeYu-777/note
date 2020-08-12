package main

import (
	"fmt"
	//"time"
)

var c chan string

func PingPong () {
	i:=0
	// goroutine
	for  {
		fmt.Println(<-c)
		c<-fmt.Sprintf("From PingPong: Hi #%d", i)
		i ++
	}
}

func main () {
	c = make(chan string)
	go PingPong()
	for i:=0;  i<10; i++ {
		c<-fmt.Sprintf("From PingPong: Hello #%d", i)
		fmt.Println(<-c)
	}

	//c := make(chan bool)
	//go func() {
	//	//fmt.Println("GO GO GO!!")
	//	//c <- true // 存channel
	//	//// 关闭通道channel
	//	//close(c)
	//	for v:=range c {
	//		fmt.Println(v)
	//	}
	//}()
	//for {
	//c := make(chan bool)
	//// select 带有超时执行什么 如下
	//select {
	//case v:= <-c:
	//	fmt.Println(v)
	//case <-time.After(3 * time.Second):
	//	fmt.Println("timeout")
	//}
	//}

	//<-c // 取channel

	// 在迭代channel的时候一定要明确关闭它close(c)
	//for v:= range c {
	//	fmt.Println(v)
	//}

}

