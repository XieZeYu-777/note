package main

import (
	"fmt"
	"sync"
	"time"
)

// todo channel
var chanInt chan int = make(chan int, 10)
var timeout chan bool = make(chan bool, 10)

var WG sync.WaitGroup // todo 协程同步


func main()  {
	// todo 并发的实现 协程 比线程小方便 | 多协程通信 | 多协程同步
	// 并发
	// 获取cpu个数
	// todo 多核cpu设置
	//fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()), "cpu")
	// 通常情况下减1 程序跑3个核心
	//runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	// todo 启动协程
	//go goConcurrency()
	//go goConcurrency()
	// 延迟
	//time.Sleep(time.Second * 60)
	// todo 协程通信
	// todo channel
	// 启动发送Send协程
	//go Send()
	//// 启动接收Receive协程
	//go Receive()
	//// 延迟执行
	//time.Sleep(time.Second * 60)
	Read()
	go Write()
	WG.Wait()
	// todo 协程同步 用到golang的系统包sync.waitgroup
	// todo Add(delta int) 添加协程记录
	// todo Done() 移出协程记录
	// todo Wait() 同步等待所有记录的协程全部结束
}

// 读取数据
func Read () {
	for i := 0; i < 3; i++  {
		// 添加一条记录
		WG.Add(1)
	}
}

// 写入数据
func Write () {
	for i := 0; i < 3; i++  {
		fmt.Println("移出", i)
		time.Sleep(time.Second * 4)
		// 完成添加记录 并移除
		WG.Done()
	}
}

// 发送数据
func Send () {
	// todo channel
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true

}

// 接收数据
func Receive () {
	//num := <- chanInt
	//fmt.Printf("接收的num值%d", num)
	//num = <- chanInt
	//fmt.Printf("接收的num值%d", num)
	//num = <- chanInt
	//fmt.Printf("接收的num值%d", num)
	// todo select
	for {
		select {
		case num:= <-chanInt:
			fmt.Println("num值是：",num)
		case <-timeout:
			fmt.Println("loading...")
		}
	}
}

func goConcurrency()  {
	for i := 1; i < 10; i++ {
		fmt.Printf("%d,", i)
	}
}
